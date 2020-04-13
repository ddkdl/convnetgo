package dnn

import "errors"

//Tensor is the basic data structure of convolutional neural networks.  It can be used for regular neural networks too.
type Tensor struct {
	dims    []int     // I use int because there is a ton of libraries in c, and this will make the conversion easier.
	stride  []int     //stride is used for placement of values and grabing values.
	f32data []float32 //place to hold data in float32
	nhwc    bool
}

const maxdimsize = 4
const mindimsize = 4

//CreateTensor creates a tensor according to the values passed. If len(dims) not ==4 an error will return!
//f64 is a place holder only thing available is float32
func CreateTensor(dims []int, NHWC bool) (*Tensor, error) {
	if len(dims) < 4 || len(dims) > 4 {
		return nil, errors.New("Not a valid length of dims")
	}

	return &Tensor{
		dims:    dims,
		stride:  findstride(dims),
		f32data: make([]float32, findvolume(dims)),
		nhwc:    NHWC,
	}, nil
}

func (t *Tensor) Set(value float32, dimlocation []int) {
	var location int
	for i := range dimlocation {
		location += dimlocation[i] * t.stride[i]
	}
	t.f32data[location] = value
}
func (t *Tensor) Get(dimlcoation []int) (value float32) {
	var location int
	for i := range dimlcoation {
		location += dimlcoation[i] * t.stride[i]
	}
	value = t.f32data[location]
	return value
}

//SetAll sets all the elments in t to value.
func (t *Tensor) SetAll(value float32) {
	for i := range t.f32data {
		t.f32data[i] = value
	}
}

//LoadFromSlice will load from values into the tensor. It starts at zero til the length of values.
//If values is longer than the volume of the tensor an error will return.
func (t *Tensor) LoadFromSlice(values []float32) (err error) {
	if len(values) > len(t.f32data) {
		return errors.New("values slice larger than tensor volume")
	}
	for i := range values {
		t.f32data[i] = values[i]
	}
	return nil
}
func findrecursivelocation(dims, stride []int, val int) (location int) {
	if len(dims) == 1 {
		val += dims[0] * stride[0]
		return val

	}
	val += dims[0] * stride[0]
	return findrecursivelocation(dims[1:], stride[1:], val)
}
func comparedims(xdims, ydims []int) bool {
	if len(xdims) != len(ydims) {
		return false
	}
	for i := range xdims {
		if xdims[i] != ydims[i] {
			return false
		}

	}
	return true
}

//findvolume will save us a lot of trouble.
func findvolume(dims []int) (vol int) {
	vol = 1
	for i := range dims {
		vol *= dims[i]
	}
	return vol
}
func findstride(dims []int) (stride []int) {
	stride = make([]int, len(dims))
	strider := int(1)
	for i := len(dims) - 1; i >= 0; i-- {
		stride[i] = strider
		strider *= dims[i]
	}
	return stride
}
