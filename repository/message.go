package repository

type Message struct {
	ID        int
	SenderID  int
	ChannelID int
	Body      string
}

type MessageFilterOptions struct {
	ChannelID int
	top       int
	skip      int
}

type MessageRepository interface {
	GetByID(id int) (*Message, error)
	GetAll(opts *MessageFilterOptions) ([]*Message, error)
	Create(message *Message) error
	Update(message *Message) error
	Delete(message *Message) error
}

type messageRepository struct {
	messages []*Message
}

// Create implements MessageRepository
func (*messageRepository) Create(message *Message) error {
	panic("unimplemented")
}

// Delete implements MessageRepository
func (*messageRepository) Delete(message *Message) error {
	panic("unimplemented")
}

// GetAll implements MessageRepository
func (*messageRepository) GetAll() ([]*Message, error) {
	panic("unimplemented")
}

// GetByID implements MessageRepository
func (*messageRepository) GetByID(id int) (*Message, error) {
	panic("unimplemented")
}

// Update implements MessageRepository
func (*messageRepository) Update(message *Message) error {
	panic("unimplemented")
}

func NewMessageRepository() MessageRepository {
	return &messageRepository{
		messages: []*Message{},
	}
}
