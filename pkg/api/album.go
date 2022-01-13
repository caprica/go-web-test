package api

type AlbumService interface {
	ArtistAlbumNames(albumArtistName string) ([]string, error)
}

type AlbumRepository interface {
	FindAlbums(albumArtistName string) ([]Album, error)
}

type albumService struct {
	storage AlbumRepository
}

func NewAlbumService(albumRepository AlbumRepository) AlbumService {
	return &albumService{storage: albumRepository}
}

// ArtistAlbumNames returns the list of album names known for the given artist name.
func (a *albumService) ArtistAlbumNames(albumArtistName string) ([]string, error) {
	albums, err := a.storage.FindAlbums(albumArtistName)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, album := range albums {
		result = append(result, album.MediaId.AlbumName)
	}
	return result, nil
}
