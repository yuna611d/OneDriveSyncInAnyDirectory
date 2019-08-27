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
	Audio struct {
		OdataType string `json:"@odata.type"`
	} `json:"audio"`
	Content struct {
		OdataType string `json:"@odata.type"`
	} `json:"content"`
	CTag    string `json:"cTag"`
	Deleted struct {
		OdataType string `json:"@odata.type"`
	} `json:"deleted"`
	Description string `json:"description"`
	File        struct {
		OdataType string `json:"@odata.type"`
	} `json:"file"`
	FileSystemInfo struct {
		OdataType string `json:"@odata.type"`
	} `json:"fileSystemInfo"`
	Folder struct {
		OdataType string `json:"@odata.type"`
	} `json:"folder"`
	Image struct {
		OdataType string `json:"@odata.type"`
	} `json:"image"`
	Location struct {
		OdataType string `json:"@odata.type"`
	} `json:"location"`
	Package struct {
		OdataType string `json:"@odata.type"`
	} `json:"package"`
	Photo struct {
		OdataType string `json:"@odata.type"`
	} `json:"photo"`
	Publication struct {
		OdataType string `json:"@odata.type"`
	} `json:"publication"`
	RemoteItem struct {
		OdataType string `json:"@odata.type"`
	} `json:"remoteItem"`
	Root struct {
		OdataType string `json:"@odata.type"`
	} `json:"root"`
	SearchResult struct {
		OdataType string `json:"@odata.type"`
	} `json:"searchResult"`
	Shared struct {
		OdataType string `json:"@odata.type"`
	} `json:"shared"`
	SharepointIds struct {
		OdataType string `json:"@odata.type"`
	} `json:"sharepointIds"`
	Size          int `json:"size"`
	SpecialFolder struct {
		OdataType string `json:"@odata.type"`
	} `json:"specialFolder"`
	Video struct {
		OdataType string `json:"@odata.type"`
	} `json:"video"`
	WebDavURL  string `json:"webDavUrl"`

	// relationships
	Activities []struct {
		OdataType string `json:"@odata.type"`
	} `json:"activities"`
	Analytics struct {
		OdataType string `json:"@odata.type"`
	} `json:"analytics"`
	Children []struct {
		OdataType string `json:"@odata.type"`
	} `json:"children"`
	CreatedByUser struct {
		OdataType string `json:"@odata.type"`
	} `json:"createdByUser"`
	LastModifiedByUser struct {
		OdataType string `json:"@odata.type"`
	} `json:"lastModifiedByUser"`
	Permissions []struct {
		OdataType string `json:"@odata.type"`
	} `json:"permissions"`
	Subscriptions []struct {
		OdataType string `json:"@odata.type"`
	} `json:"subscriptions"`
	Thumbnails []struct {
		OdataType string `json:"@odata.type"`
	} `json:"thumbnails"`
	Versions []struct {
		OdataType string `json:"@odata.type"`
	} `json:"versions"`

	// inherited from baseItem
	ID        string `json:"id"`
	CreatedBy struct {
		OdataType string `json:"@odata.type"`
	} `json:"createdBy"`
	CreatedDateTime string `json:"createdDateTime"`
	ETag            string `json:"eTag"`
	LastModifiedBy  struct {
		OdataType string `json:"@odata.type"`
	} `json:"lastModifiedBy"`
	LastModifiedDateTime string `json:"lastModifiedDateTime"`
	Name                 string `json:"name"`
	ParentReference      struct {
		OdataType string `json:"@odata.type"`
	} `json:"parentReference"`
	WebURL                         string `json:"webUrl"`
	// instance annotations
	MicrosoftGraphConflictBehavior string `json:"@microsoft.graph.conflictBehavior"`
	MicrosoftGraphDownloadURL      string `json:"@microsoft.graph.downloadUrl"`
	MicrosoftGraphSourceURL        string `json:"@microsoft.graph.sourceUrl"`
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
