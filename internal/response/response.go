package response

type AirbnbResponse struct {
    Results []struct {
        Position       int             `json:"position"`
        Name           string          `json:"name"`
        Bathrooms      float32             `json:"bathrooms"`
        Beds           int             `json:"beds"`
        City           string          `json:"city"`
        Neighborhood   string          `json:"neighborhood"`
        ReviewsCount   int             `json:"reviewsCount"`
        Rating         float32         `json:"rating"`
        Type           string          `json:"type"`
        Address        string          `json:"address"`
        Facilities     []string        `json:"previewAmenities"`
        Price          struct {
            Currency   string          `json:"currency"`
            Total      int             `json:"total"`
            PriceItems []struct {
                Title  string          `json:"title"`
                Amount int             `json:"amount"`
            }                           `json:"priceItems"`
        }                               `json:"price"`
    }                                   `json:"results"`
}
