package ringbuffer

type RingBuffer struct {
	frontPos int32
	rearPos  int32

	defaultSize int32

	Buffer []byte
}

func NewRingBuffer(size int32) *RingBuffer {
	return &RingBuffer{
		frontPos:    0,
		rearPos:     0,
		Buffer:      make([]byte, 0, size),
		defaultSize: size,
	}
}

func (c *RingBuffer) Enqueue(data *[]byte, size int32) int32 {
	var tmpRearPos int32 = c.rearPos
	var tmpFrontPos int32 = c.frontPos
	var ret_val int32 = 0
	for size > 0 {
		if (tmpRearPos + 1%c.defaultSize) == tmpFrontPos {
			break
		}

		//(*c.Buffer)[tmpRearPos] = *data[]
		c.Buffer[tmpRearPos] = (*data)[ret_val]

		ret_val++
		size--
	}
	c.rearPos = tmpRearPos
	return ret_val
}
