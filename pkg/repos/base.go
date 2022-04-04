package repos

type BaseEntity struct {
	Id ID `json:"id" bson:"id"`
}

func (c *BaseEntity) GetId() ID {
	return c.Id
}

func (c *BaseEntity) SetId(id ID) Entity {
	c.Id = id
	return c
}
