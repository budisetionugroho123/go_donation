package dto

type OrganizationRequest struct {
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	OrganizationName string `json:"organization_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	RoleID           int    `json:"role_id"`
	Description      string `json:"description"`
	LogoURL          string `json:"logo_url"`
	ContactInfo      string `json:"contact_info"`
}

type OrganizationResponse struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	OrganizationID   uint   `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	RoleID           int    `json:"role_id"`
	Description      string `json:"description"`
	LogoURL          string `json:"logo_url"`
	ContactInfo      string `json:"contact_info"`
}
