package yamlutil

import (
	"crypto/sha1"
	"encoding/base64"
	"hash"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
)

// Hash computes a hash of the given node
func Hash(node *yaml.Node) string {
	hash := sha1.New()
	hashNext(hash, node)

	sum := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(sum)
}

func hashNext(hash hash.Hash, node *yaml.Node) {
	switch node.Kind {
	case yaml.DocumentNode:
		hashNext(hash, node.Content[0])
	case yaml.SequenceNode:
		for _, child := range node.Content {
			hashNext(hash, child)
		}
	case yaml.MappingNode:
		for i := 0; (i + 1) < len(node.Content); i += 2 {
			hashNext(hash, node.Content[i])
			hashNext(hash, node.Content[i+1])
		}
	case yaml.ScalarNode:
		switch node.Tag {
		case "!!int", "!!float":
			d, err := decimal.NewFromString(node.Value)
			if err != nil {
				hash.Write([]byte(node.Value))
			} else {
				// Normalize ints and floats to a consistent format
				hash.Write([]byte(d.String()))
			}
		default:
			hash.Write([]byte(node.Value))
		}
	case yaml.AliasNode:
		hashNext(hash, node.Alias)
	}
}
