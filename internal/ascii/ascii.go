package ascii

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/qeesung/image2ascii/convert"
	"sonarr-sabnzbd-cli/internal/models"
)

// Config holds ASCII art conversion settings
type Config struct {
	Width     int
	Height    int
	Colored   bool
	CacheDir  string
	ImageType string // "poster", "banner", "fanart"
}

// DefaultConfig returns a default ASCII configuration
func DefaultConfig() *Config {
	cacheDir := filepath.Join(os.TempDir(), "sonarr-cli-ascii-cache")
	os.MkdirAll(cacheDir, 0755)
	return &Config{
		Width:     8,    // Small 8x8 like scalarr
		Height:    8,    // Small 8x8 like scalarr
		Colored:   true, // Enable colors like scalarr
		CacheDir:  cacheDir,
		ImageType: "poster",
	}
}

// ConvertImageToASCII converts an image URL to ASCII art
func ConvertImageToASCII(url string, config *Config) (string, error) {
	if config == nil {
		config = DefaultConfig()
	}

	// Try to get from cache first
	cacheKey := strings.ReplaceAll(strings.ReplaceAll(url, "/", "_"), ":", "_")
	cachePath := filepath.Join(config.CacheDir, cacheKey+".ascii")

	if ascii, err := readFromCache(cachePath); err == nil {
		return ascii, nil
	}

	// Download and convert image
	img, err := downloadImage(url)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %w", err)
	}

	// Resize image maintaining aspect ratio
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Calculate new dimensions maintaining aspect ratio
	aspectRatio := float64(width) / float64(height)
	if config.Width > 0 && config.Height > 0 {
		// Both dimensions specified
		img = imaging.Resize(img, config.Width, config.Height, imaging.Lanczos)
	} else if config.Width > 0 {
		// Only width specified, calculate height
		newHeight := int(float64(config.Width) / aspectRatio * 0.5) // 0.5 to account for character aspect ratio
		img = imaging.Resize(img, config.Width, newHeight, imaging.Lanczos)
	} else if config.Height > 0 {
		// Only height specified, calculate width
		newWidth := int(float64(config.Height) * aspectRatio * 2) // 2 to account for character aspect ratio
		img = imaging.Resize(img, newWidth, config.Height, imaging.Lanczos)
	}

	// Convert to ASCII with improved options
	converter := convert.NewImageConverter()
	options := &convert.Options{
		FixedWidth:  config.Width,
		FixedHeight: config.Height,
		Colored:     config.Colored, // Use config setting for colors
		Reversed:    false,
	}

	ascii := converter.Image2ASCIIString(img, options)

	// Only strip ANSI codes if colors are disabled
	if !config.Colored {
		ascii = stripANSI(ascii)
	}

	// Format the ASCII art into proper lines for 8x8 display
	lines := make([]string, 0, config.Height)
	for i := 0; i < config.Height && i*config.Width < len(ascii); i++ {
		start := i * config.Width
		end := start + config.Width
		if end > len(ascii) {
			end = len(ascii)
		}
		lines = append(lines, ascii[start:end])
	}
	formattedASCII := strings.Join(lines, "\n")

	// Cache the result
	if err := writeToCache(cachePath, formattedASCII); err != nil {
		// Don't fail if caching fails, just log
		fmt.Fprintf(os.Stderr, "Warning: failed to cache ASCII art: %v\n", err)
	}

	return formattedASCII, nil
}

// GetSeriesPosterASCII converts a series poster to ASCII art
func GetSeriesPosterASCII(series models.Series, config *Config) (string, error) {
	if config == nil {
		config = DefaultConfig()
	}

	// Find the best image for the requested type
	imageURL := findBestImage(series.Images, config.ImageType)
	if imageURL == "" {
		return "", fmt.Errorf("no suitable image found for series %s", series.Title)
	}

	return ConvertImageToASCII(imageURL, config)
}

// findBestImage finds the best image URL for the given type
func findBestImage(images []models.SeriesImage, preferredType string) string {
	// First try to find the preferred type
	for _, img := range images {
		if strings.EqualFold(img.CoverType, preferredType) {
			return img.URL
		}
	}

	// Fallback to poster if preferred type not found
	for _, img := range images {
		if strings.EqualFold(img.CoverType, "poster") {
			return img.URL
		}
	}

	// Last resort: return first available image
	if len(images) > 0 {
		return images[0].URL
	}

	return ""
}

// downloadImage downloads an image from URL
func downloadImage(url string) (image.Image, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	img, _, err := image.Decode(resp.Body)
	return img, err
}

// readFromCache reads ASCII art from cache
func readFromCache(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// writeToCache writes ASCII art to cache
func writeToCache(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// ClearCache removes all cached ASCII art
func ClearCache(config *Config) error {
	if config == nil {
		config = DefaultConfig()
	}
	return os.RemoveAll(config.CacheDir)
}

// stripANSI removes ANSI escape codes from a string
func stripANSI(str string) string {
	// Simple regex to match ANSI color codes in the format [38;5;16m
	ansiRegex := regexp.MustCompile(`\[[0-9;]*m`)
	return ansiRegex.ReplaceAllString(str, "")
}

// FormatASCIIWithTitle formats ASCII art with a title
func FormatASCIIWithTitle(title, ascii string) string {
	lines := strings.Split(ascii, "\n")
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Create a border
	border := strings.Repeat("─", maxWidth+4)
	titleLine := fmt.Sprintf("┌%s┐", border)
	contentLine := fmt.Sprintf("│ %-*s │", maxWidth, title)

	result := []string{titleLine, contentLine, fmt.Sprintf("├%s┤", border)}

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			result = append(result, fmt.Sprintf("│ %-*s │", maxWidth, line))
		}
	}

	result = append(result, fmt.Sprintf("└%s┘", border))
	return strings.Join(result, "\n")
}
