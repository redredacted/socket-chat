package repository

type Channel struct {
	ID    int
	Title string
}

type ChannelRepository interface {
	GetByID(id int) (*Channel, error)
	GetAll() ([]*Channel, error)
	Create(channel *Channel) error
	Update(channel *Channel) error
	Delete(channel *Channel) error
}

type channelRepository struct {
	channels []*Channel
}

func NewChannelRepository() ChannelRepository {
	return &channelRepository{
		channels: []*Channel{},
	}
}

// Create implements ChannelRepository
func (*channelRepository) Create(channel *Channel) error {
	panic("unimplemented")
}

// Delete implements ChannelRepository
func (*channelRepository) Delete(channel *Channel) error {
	panic("unimplemented")
}

// GetAll implements ChannelRepository
func (*channelRepository) GetAll() ([]*Channel, error) {
	panic("unimplemented")
}

// GetByID implements ChannelRepository
func (*channelRepository) GetByID(id int) (*Channel, error) {
	panic("unimplemented")
}

// Update implements ChannelRepository
func (*channelRepository) Update(channel *Channel) error {
	panic("unimplemented")
}
