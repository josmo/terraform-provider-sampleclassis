package classis

type Login struct {
	EmailAddress string `json:"email"`
	Password     string `json:"password"`
}

type LoginResponse struct {
	ID           string `json:"_id,omitempty"`
	Token        string `json:"token,omitempty"`
	TokenExpires string `json:"tokenExpires,omitempty"`
}

type SpotGroup struct {
	ID                  string              `json:"_id,omitempty"`
	Name                string              `json:"name,omitempty"`
	Region              string              `json:"region",omitempty`
	IamFleetRole        string              `json:"iamFleetRole,omitempty"`
	DesiredQty          string              `json:"desiredQty,omitempty"`
	Quantity            string              `json:"quantity,omitempty"`
	Vpc                 string              `json:"vpc,omitempty"`
	LaunchSpecification LaunchSpecification `json:"launchSpecification,omitempty"`
	TypesSelected       []string            `json:"typesSelected,omitempty"`
}

type LaunchSpecification struct {
	ImageId  string `json:"imageId,omitempty"`
	SubnetId string `json:"subnetId,omitempty"`
	KeyName  string `json:"keyName,omitempty"`
	DefaultDeviceSize string `json:"defaultDeviceSize,omitempty"`
	SecurityGroups []SecurityGroup `json:"securityGroups,omitempty"`
}

type SecurityGroup struct {
	GroupId string `json:"groupId,omitempty"`
}
