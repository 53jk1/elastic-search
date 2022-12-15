package domain

type SearchRecord struct {
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
}

func NewSearchRecord() *SearchRecord {
	return &SearchRecord{}
}

func (s *SearchRecord) SetID(id string) {
	s.ID = id
}

func (s *SearchRecord) SetTitle(title string) {
	s.Title = title
}

func (s *SearchRecord) SetDescription(description string) {
	s.Description = description
}

func (s *SearchRecord) SetURL(url string) {
	s.URL = url
}

func (s *SearchRecord) SetImageURL(imageURL string) {
	s.ImageURL = imageURL
}

func (s *SearchRecord) SetCreatedAt(createdAt string) {
	s.CreatedAt = createdAt
}

func (s *SearchRecord) SetUpdatedAt(updatedAt string) {
	s.UpdatedAt = updatedAt
}

func (s *SearchRecord) GetID() string {
	return s.ID
}

func (s *SearchRecord) GetTitle() string {
	return s.Title
}

func (s *SearchRecord) GetDescription() string {
	return s.Description
}

func (s *SearchRecord) GetURL() string {
	return s.URL
}

func (s *SearchRecord) GetImageURL() string {
	return s.ImageURL
}

func (s *SearchRecord) GetCreatedAt() string {
	return s.CreatedAt
}

func (s *SearchRecord) GetUpdatedAt() string {
	return s.UpdatedAt
}

func (s *SearchRecord) SetRecord(record *Record) {
	s.SetID(record.GetID())
	s.SetTitle(record.GetTitle())
	s.SetDescription(record.GetDescription())
	s.SetURL(record.GetURL())
	s.SetImageURL(record.GetImageURL())
	s.SetCreatedAt(record.GetCreatedAt())
	s.SetUpdatedAt(record.GetUpdatedAt())
}

func (s *SearchRecord) GetRecord() *Record {
	record := NewRecord()
	record.SetID(s.GetID())
	record.SetTitle(s.GetTitle())
	record.SetDescription(s.GetDescription())
	record.SetURL(s.GetURL())
	record.SetImageURL(s.GetImageURL())
	record.SetCreatedAt(s.GetCreatedAt())
	record.SetUpdatedAt(s.GetUpdatedAt())
	return record
}

func (s *SearchRecord) SetSearchRecord(searchRecord *SearchRecord) {
	s.SetID(searchRecord.GetID())
	s.SetTitle(searchRecord.GetTitle())
	s.SetDescription(searchRecord.GetDescription())
	s.SetURL(searchRecord.GetURL())
	s.SetImageURL(searchRecord.GetImageURL())
	s.SetCreatedAt(searchRecord.GetCreatedAt())
	s.SetUpdatedAt(searchRecord.GetUpdatedAt())
}

func (s *SearchRecord) GetSearchRecord() *SearchRecord {
	searchRecord := NewSearchRecord()
	searchRecord.SetID(s.GetID())
	searchRecord.SetTitle(s.GetTitle())
	searchRecord.SetDescription(s.GetDescription())
	searchRecord.SetURL(s.GetURL())
	searchRecord.SetImageURL(s.GetImageURL())
	searchRecord.SetCreatedAt(s.GetCreatedAt())
	searchRecord.SetUpdatedAt(s.GetUpdatedAt())
	return searchRecord
}

func (s *SearchRecord) SetSearchRecords(searchRecords []*SearchRecord) {
	for _, searchRecord := range searchRecords {
		s.SetSearchRecord(searchRecord)
	}
}

func (s *SearchRecord) GetSearchRecords() []*SearchRecord {
	var searchRecords []*SearchRecord
	for _, searchRecord := range searchRecords {
		searchRecords = append(searchRecords, searchRecord.GetSearchRecord())
	}
	return searchRecords
}
