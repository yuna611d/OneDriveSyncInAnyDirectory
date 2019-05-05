package models

type OneDriveIdentity struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}
type OneDriveIdentitySet struct {
	Application OneDriveIdentity `json:"application"`
	Device      OneDriveIdentity `json:"device"`
	Group       OneDriveIdentity `json:"group"`
	User        OneDriveIdentity `json:"user"`
}

type OneDriveItemReference struct {
	DriveID   string `json:"driveId"`
	DriveType string `json:"driveType"`
	ID        string `json:"id"`
	ListID    string `json:"listId"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	ShareID   string `json:"shareId"`
	SiteID    string `json:"siteId"`
}

type OneDriveBaseItem struct {
	ID                   string                `json:"id"`
	CreatedBy            OneDriveIdentitySet   `json:"createdBy"`
	CreatedDateTime      string                `json:"createdDateTime"`
	Description          string                `json:"description"`
	ETag                 string                `json:"eTag"`
	LastModifiedBy       OneDriveIdentitySet   `json:"lastModifiedBy"`
	LastModifiedDateTime string                `json:"lastModifiedDateTime"`
	Name                 string                `json:"name"`
	ParentReference      OneDriveItemReference `json:"parentReference"`
	WebURL               string                `json:"webUrl"`
}
type OneDriveDriveItems struct {
	Value []OneDriveDriveItem `json:"value"`
}
type OneDriveDriveItem struct {
	/* inherited from baseItem */
	DriveID   string `json:"driveId"`
	DriveType string `json:"driveType"`
	ID        string `json:"id"`
	ListID    string `json:"listId"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	ShareID   string `json:"shareId"`
	SiteID    string `json:"siteId"`
	/* */
	Folder OneDriveFolder `json:"folder"`
	/* */
	File OneDriveFile `json:"file"`
}
type OneDriveFolder struct {
	ChildCount int `json:"childCount"`
}
type OneDriveFile struct {
	Hashes   OneDriveHashes `json:"hashes"`
	MimeType string         `json:"mimeType"`
}

type OneDriveHashes struct {
	Crc32Hash    string `json:"crc32Hash"`
	Sha1Hash     string `json:"sha1Hash"`
	QuickXorHash string `json:"quickXorHash"`
}
