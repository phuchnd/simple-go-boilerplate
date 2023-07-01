package generators

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
)

type SnowflakeID snowflake.ID

//go:generate mockery --name=ISnowflakeIDGenerator --case=snake --disable-version-string
type ISnowflakeIDGenerator interface {
	// Next generates a new SnowflakeID.
	Next() SnowflakeID
}

type snowflakeIDGeneratorImpl struct {
	node *snowflake.Node
}

func NewSnowflakeIDGenerator() (ISnowflakeIDGenerator, error) {
	node, err := snowflake.NewNode(int64(rand.Int31n(1024)))
	if err != nil {
		return nil, err
	}

	return &snowflakeIDGeneratorImpl{node}, nil
}

func (impl *snowflakeIDGeneratorImpl) Next() SnowflakeID {
	return SnowflakeID(impl.node.Generate())
}
