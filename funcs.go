package validator

import (
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/inhies/go-bytesize"
)

// IsUsername check if string is valid username
func IsUsername(username string) bool {
	r := regexp.MustCompile(`^[0-9a-zA-Z\-\._]+$`)
	return r.MatchString(username)
}

// IsTel check if string is valid tel (IR numbers)
func IsTel(tel string) bool {
	rx := regexp.MustCompile(`[\d]+`)
	rm := regexp.MustCompile(`^0\d{10}$`)
	return rm.MatchString(strings.Join(rx.FindAllString(tel, -1), ""))
}

// IsMobile check if string is valid mobile number (IR numbers)
func IsMobile(mobile string) bool {
	rx := regexp.MustCompile(`[\d]+`)
	rm := regexp.MustCompile(`^09\d{9}$`)
	return rm.MatchString(strings.Join(rx.FindAllString(mobile, -1), ""))
}

// IsPostalcode check if string is valid postal code
func IsPostalcode(postalCode string) bool {
	rx := regexp.MustCompile(`[\d]+`)
	rm := regexp.MustCompile(`^\d{10}$`)
	return rm.MatchString(strings.Join(rx.FindAllString(postalCode, -1), ""))
}

// IsIdentifier check if string is valid identifier
func IsIdentifier(id string) bool {
	idf, err := strconv.Atoi(id)
	return err == nil && idf > 0
}

// IsUnsigned check if string is unsigned number
func IsUnsigned(num string) bool {
	idf, err := strconv.Atoi(num)
	return err == nil && idf >= 0
}

// IsIDNumber check if string is valid id number
func IsIDNumber(idNum string) bool {
	r := regexp.MustCompile(`^\d{1,10}$`)
	return r.MatchString(idNum)
}

// ISNationalCode check if string is valid id national code
func ISNationalCode(idNum string) bool {
	rx := regexp.MustCompile(`[\d]+`)
	rm := regexp.MustCompile(`^\d{10}$`)
	return rm.MatchString(strings.Join(rx.FindAllString(idNum, -1), ""))
}

// IsCreditCardNumber check if string is valid id credit card number
func IsCreditCardNumber(num string) bool {
	rx := regexp.MustCompile(`[\d]+`)
	rm := regexp.MustCompile(`^(\d{16})|(\d{20})$`)
	return rm.MatchString(strings.Join(rx.FindAllString(num, -1), ""))
}

// IsUUID check if string is valid uuid
func IsUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// IsIP check if address if a valid ip
func IsIP(address string) bool {
	r := regexp.MustCompile(`^(([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])\.){3}([0–9]|[1–9][0–9]|1[0–9]{2}|2[0–4][0–9]|25[0–5])$`)
	return r.MatchString(address)
}

// IsIPPort check if address if a valid ip contains port
func IsIPPort(address string) bool {
	r := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):[0-9]+$`)
	return r.MatchString(address)
}

// ValidateUploadSize check if file size in range
// use B, KB, MB, GB for size string
// ex: 1KB, 3MB
// Note: not use float point!
func ValidateUploadSize(file *multipart.FileHeader, min string, max string) (bool, error) {
	minSize, err := bytesize.Parse(min)
	if err != nil {
		return false, err
	}

	maxSize, err := bytesize.Parse(max)
	if err != nil {
		return false, err
	}

	return (uint64(file.Size) >= uint64(minSize) && uint64(file.Size) <= uint64(maxSize)), nil
}

// ValidateUploadMime check if file upload mime is valid
func ValidateUploadMime(file *multipart.FileHeader, mimes ...string) (bool, error) {
	f, err := file.Open()
	if err != nil {
		return false, err
	}
	defer f.Close()

	if mime, err := mimetype.DetectReader(f); err != nil {
		return false, err
	} else if mime != nil {
		return mimetype.EqualsAny(mime.String(), mimes...), nil
	}

	return false, nil
}

// ValidateUploadExt check if file upload extension is valid
func ValidateUploadExt(file *multipart.FileHeader, exts ...string) (bool, error) {
	f, err := file.Open()
	if err != nil {
		return false, err
	}
	defer f.Close()

	if mime, err := mimetype.DetectReader(f); err != nil {
		return false, err
	} else if mime != nil {
		for _, ext := range exts {
			if strings.ToLower(ext) == strings.ToLower(mime.Extension()) {
				return true, nil
			}
		}
	}

	return false, nil
}
