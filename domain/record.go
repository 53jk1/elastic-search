package domain

type Record struct {
	// ID is the unique identifier of the record
	ID string `json:"id"`
	// Title is the title of the record
	Title string `json:"title"`
	// Description is the description of the record
	Description string `json:"description"`
	// URL is the URL of the record
	URL string `json:"url"`
	// ImageURL is the URL of the image of the record
	ImageURL string `json:"image_url"`
	// CreatedAt is the date and time when the record was created
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the date and time when the record was updated
	UpdatedAt string `json:"updated_at"`
	// FullSearchField is the field that is used for full text search
	FullSearchField string `json:"full_search_field"`
}

func NewRecord() *Record {
	return &Record{}
}

func (r *Record) SetID(id string) {
	r.ID = id
}

func (r *Record) SetTitle(title string) {
	r.Title = title
}

func (r *Record) SetDescription(description string) {
	r.Description = description
}

func (r *Record) SetURL(url string) {
	r.URL = url
}

func (r *Record) SetImageURL(imageURL string) {
	r.ImageURL = imageURL
}

func (r *Record) SetCreatedAt(createdAt string) {
	r.CreatedAt = createdAt
}

func (r *Record) SetUpdatedAt(updatedAt string) {
	r.UpdatedAt = updatedAt
}

func (r *Record) GetID() string {
	return r.ID
}

func (r *Record) GetTitle() string {
	return r.Title
}

func (r *Record) GetDescription() string {
	return r.Description
}

func (r *Record) GetURL() string {
	return r.URL
}

func (r *Record) GetImageURL() string {
	return r.ImageURL
}

func (r *Record) GetCreatedAt() string {
	return r.CreatedAt
}

func (r *Record) GetUpdatedAt() string {
	return r.UpdatedAt
}

func SetFullSearchField(record *Record) {
	record.SetFullSearchField(record.GetTitle() + " " + record.GetDescription())
}

func (r *Record) SetFullSearchField(fullSearchField string) {
	r.FullSearchField = fullSearchField
}

func (r *Record) GetFullSearchField() string {
	return r.FullSearchField
}
