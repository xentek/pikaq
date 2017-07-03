package pikaq

import (
	"fmt"

	"github.com/google/uuid"
)

type ConsumerTag struct {
	name string
	id   uuid.UUID
}

func (c ConsumerTag) Tag() string {
	return fmt.Sprintf("%s_%s", c.name, c.id.String())
}

func NewConsumerTag(name string) *ConsumerTag {
	return &ConsumerTag{name: name, id: uuid.New()}
}
