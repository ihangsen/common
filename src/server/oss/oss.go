package oss

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	openapicred "github.com/aliyun/credentials-go/credentials"
	"github.com/google/uuid"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/collection/dict"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/i18n"
	"github.com/ihangsen/common/src/utils/strs"
	"github.com/ihangsen/common/src/utils/trans"
	"sync"
	"time"
)

type AliSecret struct {
	AccessKey string
	SecretKey string
}

type OssConf struct {
	//所处地域
	Region string
	//桶名
	BucketName string
	//阿里云产品名
	Product string
	//OSS上传地址
	AliSecret     *AliSecret
	RoleType      string
	UploadTypeMap dict.Dict[uint8, string]
}

var (
	ossClient = new(oss.Client)
	ossConf   = new(OssConf)
	once      sync.Once
)

const privateBound = 100

func Init(conf *OssConf) {
	once.Do(func() {
		ossConf = conf
		secret := ossConf.AliSecret
		config := new(openapicred.Config).
			SetType(ossConf.RoleType).
			SetAccessKeyId(secret.AccessKey).
			SetAccessKeySecret(secret.SecretKey)
		credential, err0 := openapicred.NewCredential(config)
		provider := credentials.CredentialsProviderFunc(func(ctx context.Context) (credentials.Credentials, error) {
			if err0 != nil {
				return credentials.Credentials{}, err0
			}
			cred, err1 := credential.GetCredential()
			if err1 != nil {
				return credentials.Credentials{}, err1
			}
			return credentials.Credentials{
				AccessKeyID:     *cred.AccessKeyId,
				AccessKeySecret: *cred.AccessKeySecret,
				SecurityToken:   *cred.SecurityToken,
			}, nil
		})
		cfg := oss.LoadDefaultConfig().
			WithCredentialsProvider(provider).
			WithRegion(ossConf.Region)
		ossClient = oss.NewClient(cfg)
	})
}

// PutUrl 获取预签名上传的url
func PutUrl(fileType uint8, fileName string) (string, string) {
	value := ossConf.UploadTypeMap.Load(fileType).Expect(i18n.Get.UploadCategoryErr)
	now := time.Now()
	filePath := strs.Join(value, "/", trans.Int2Str(now.Year()), "/", trans.Int2Str(int(now.Month())), "/",
		trans.Int2Str(now.Day()), "/", uuid.New().String(), "-", fileName)
	aclType := oss.ObjectACLDefault
	if fileType > privateBound {
		aclType = oss.ObjectACLPrivate
	}
	result := catch.Try1(ossClient.Presign(context.Background(), &oss.PutObjectRequest{
		Bucket: oss.Ptr(ossConf.BucketName),
		Key:    oss.Ptr(filePath),
		Acl:    aclType,
	},
		oss.PresignExpires(3*time.Minute),
	))
	return filePath, result.URL
}

// License 通过文件名获取私有桶里的文件内容
func License(filePaths vec.Vec[string]) vec.Vec[string] {
	vector := vec.New[string](len(filePaths))
	for _, filePath := range filePaths {
		result := catch.Try1(ossClient.Presign(context.Background(), &oss.GetObjectRequest{
			Bucket: oss.Ptr(ossConf.BucketName),
			Key:    oss.Ptr(filePath),
		},
			oss.PresignExpires(3*time.Minute),
		))
		vector.Append(result.URL)
	}
	return vector
}
