package tuple

import "fmt"

type T1[T0 any] struct {
	V0 T0
}

type T2[T0, T1 any] struct {
	V0 T0
	V1 T1
}

type T3[T0, T1, T2 any] struct {
	V0 T0
	V1 T1
	V2 T2
}

type T4[T0, T1, T2, T3 any] struct {
	V0 T0
	V1 T1
	V2 T2
	V3 T3
}

type T5[T0, T1, T2, T3, T4 any] struct {
	V0 T0
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

type T6[T0, T1, T2, T3, T4, T5 any] struct {
	V0 T0
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

type T7[T0, T1, T2, T3, T4, T5, T6 any] struct {
	V0 T0
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
}

type T8[T0, T1, T2, T3, T4, T5, T6, T7 any] struct {
	V0 T0
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
}

type T9[T0, T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	V0 T0
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
}

func (t T1[T0]) Unravel() T0 {
	return t.V0
}

func (t T2[T0, T1]) Unravel() (T0, T1) {
	return t.V0, t.V1
}

func (t T3[T0, T1, T2]) Unravel() (T0, T1, T2) {
	return t.V0, t.V1, t.V2
}

func (t T4[T0, T1, T2, T3]) Unravel() (T0, T1, T2, T3) {
	return t.V0, t.V1, t.V2, t.V3
}

func (t T5[T0, T1, T2, T3, T4]) Unravel() (T0, T1, T2, T3, T4) {
	return t.V0, t.V1, t.V2, t.V3, t.V4
}

func (t T6[T0, T1, T2, T3, T4, T5]) Unravel() (T0, T1, T2, T3, T4, T5) {
	return t.V0, t.V1, t.V2, t.V3, t.V4, t.V5
}

func (t T7[T0, T1, T2, T3, T4, T5, T6]) Unravel() (T0, T1, T2, T3, T4, T5, T6) {
	return t.V0, t.V1, t.V2, t.V3, t.V4, t.V5, t.V6
}

func (t T8[T0, T1, T2, T3, T4, T5, T6, T7]) Unravel() (T0, T1, T2, T3, T4, T5, T6, T7) {
	return t.V0, t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7
}

func (t T9[T0, T1, T2, T3, T4, T5, T6, T7, T8]) Unravel() (T0, T1, T2, T3, T4, T5, T6, T7, T8) {
	return t.V0, t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8
}

func (t T1[T0]) String() string {
	return fmt.Sprintf("(%v)", t.V0)
}

func (t T2[T0, T1]) String() string {
	return fmt.Sprintf("(%v,%v)", t.V0, t.V1)
}

func (t T3[T0, T1, T2]) String() string {
	return fmt.Sprintf("(%v,%v,%v)", t.V0, t.V1, t.V2)
}

func (t T4[T0, T1, T2, T3]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v)", t.V0, t.V1, t.V2, t.V3)
}

func (t T5[T0, T1, T2, T3, T4]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v)", t.V0, t.V1, t.V2, t.V3, t.V4)
}

func (t T6[T0, T1, T2, T3, T4, T5]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v)", t.V0, t.V1, t.V2, t.V3, t.V4, t.V5)
}

func (t T7[T0, T1, T2, T3, T4, T5, T6]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v)", t.V0, t.V1, t.V2, t.V3, t.V4, t.V5, t.V6)
}

func (t T8[T0, T1, T2, T3, T4, T5, T6, T7]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v)", t.V0, t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7)
}

func (t T9[T0, T1, T2, T3, T4, T5, T6, T7, T8]) String() string {
	return fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v,%v)", t.V0, t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8)
}

func T2Of[T0, T1 any](v0 T0, v1 T1) T2[T0, T1] {
	return T2[T0, T1]{v0, v1}
}

func T3Of[T0, T1, T2 any](v0 T0, v1 T1, v2 T2) T3[T0, T1, T2] {
	return T3[T0, T1, T2]{v0, v1, v2}
}

func T4Of[T0, T1, T2, T3 any](v0 T0, v1 T1, v2 T2, v3 T3) T4[T0, T1, T2, T3] {
	return T4[T0, T1, T2, T3]{v0, v1, v2, v3}
}

func T5Of[T0, T1, T2, T3, T4 any](v0 T0, v1 T1, v2 T2, v3 T3, v4 T4) T5[T0, T1, T2, T3, T4] {
	return T5[T0, T1, T2, T3, T4]{v0, v1, v2, v3, v4}
}

func T6Of[T0, T1, T2, T3, T4, T5 any](v0 T0, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) T6[T0, T1, T2, T3, T4, T5] {
	return T6[T0, T1, T2, T3, T4, T5]{v0, v1, v2, v3, v4, v5}
}

func T7Of[T0, T1, T2, T3, T4, T5, T6 any](v0 T0, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) T7[T0, T1, T2, T3, T4, T5, T6] {
	return T7[T0, T1, T2, T3, T4, T5, T6]{v0, v1, v2, v3, v4, v5, v6}
}

func T8Of[T0, T1, T2, T3, T4, T5, T6, T7 any](v0 T0, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) T8[T0, T1, T2, T3, T4, T5, T6, T7] {
	return T8[T0, T1, T2, T3, T4, T5, T6, T7]{v0, v1, v2, v3, v4, v5, v6, v7}
}

func T9Of[T0, T1, T2, T3, T4, T5, T6, T7, T8 any](v0 T0, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) T9[T0, T1, T2, T3, T4, T5, T6, T7, T8] {
	return T9[T0, T1, T2, T3, T4, T5, T6, T7, T8]{v0, v1, v2, v3, v4, v5, v6, v7, v8}
}
