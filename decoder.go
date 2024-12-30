package yamlpath

import "io"

// Decoder provides a mechanism for decoding a collection of YAML nodes into
// target objects.
type Decoder struct {
	collection Collection
}

// NewDecoder creates a new Decoder that can be used to decode a collection
// of YAML nodes into target objects.
func NewDecoder(collection Collection) *Decoder {
	return &Decoder{collection: collection}
}

// Decode decodes the next YAML node in the collection into the target object,
// and then advances the decoder to the next node.
// If the collection is exhausted, this will return an [io.EOF] error.
//
// If the target object cannot be decoded, the underlying error will be
// returned, and the decoder will not advance to the next node.
func (d *Decoder) Decode(out any) error {
	if len(d.collection) == 0 {
		return io.EOF
	}
	node := d.collection[0]
	if err := node.Decode(out); err != nil {
		return err
	}
	d.collection = d.collection[1:]
	return nil
}
