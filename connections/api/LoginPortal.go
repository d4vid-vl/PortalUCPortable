package api

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func LoginPortal(user string, password string) bool {

	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}
	resp, _ := client.PostForm("https://sso.uc.cl/cas/login?service=https%3A%2F%2Fportal.uc.cl%2Fc%2Fportal%2Flogin", url.Values{
		"password":    {password},
		"username":    {user},
		"execution":   {"43b228a1-cfed-4240-a863-08ec614e27b9_ZXlKaGJHY2lPaUpJVXpVeE1pSXNJblI1Y0NJNklrcFhWQ0o5LmJjcXBFX3VLTFVxQnJNREp3bG05T0FmYU1HY0J5OWJ0elVwMVIxQUVZZDBRc1lhSzFNRklTOTJUVGZDQVhBUUx3M1h6MWxmMXdhMDdvQkREQW93MkdwaG9aNGsxdjl5MWJhZUVWXzVWV1d5cHBqZGNFMWE0UzdvVE5PcVZNVkczQUtOQ3VpZTMxTTBHaTUzMmxNa05jY0ZwMkxhMTVKTE9ZX2lJYl82UVdWQVozeWplejVMTElOTjFDYUwteXpKcUIwTWQ2blF0WFlnQk9IeXlVeHRhSlVXMW5CUEowV3FmRjZPWkFia0txTWdwVFpyaTByWGlZeVhjTW1VcVF4ZHE2dUQ1VXdNNWJ6ZTR5dHNMVGtTT3Q3c1RVRl9aMmRFRmpoY0lpd1JMZHZTMDdpMUxwVTNkZFJpYk1yZk1JV2E2TEZQTXhzSEhJWlBFM3UtX29aTEEzN3ZsQ0d5N1ZtcHIycVpxQmpPVUxqZk1ua29KeWI4emo5cnpOLUpfOV9ldGQtNVFJaWQtWUc2cHBwczRBMlo2NFZsMEhxa3lUY0kzSXVxLXJrZkxBR2VLekc1SWVJTmk1NS1GNlF5RjZMV013ZVhoYXRGZzdfVE5wX0N1YjZ6UzhFMXNnaFZ6RHItX1VwejJueHE1NGNjeFFGN0x4ak44NzRlRjZoNUJOOVMxVzh0ZUlEY0NUUVFDNW15Wl9jem1UVkQ0blE5cnlDQ3NGZmlPeE1nOGI0cHpBZWlDUTVfcjdpOVhYNi02cTVkMXdMZ1FKbFdOSERlSUx2SjJZTVZJTzFBMktQcGxNcEhQcGJlMHVzYVQ3RGhhM3VJNzRQYzJ1dEFHdTd6ZmZ3anY5ckFEQU04TzBQcTNoaDlMU2drdjFSWUozMzR3eXNIMTRHaXFDV1pETURmYjV5Rm9DV0g4MTdsYVNOZzhZQ1JiaFowaGczemNNelU5cHY1UzVqeDlwbEVWMXl3MURLNlRBTThKSmlFNkhqTG45bk11ZFp3emNaSlNGTTl1WnRkUndad2gyOGR6TUdmRFNzcHE1SGxTV19scDZhTWFRbFlFQlo4TzBCRk9PMEt6OVExVDlFQjkyTmZRZm5iT1Y4eG5pVGo3Zy1Wc01TVWwtYU1UMGdDNVFSbGg2TDdYUW40Y0VKRW5NTHRHNEh3SXZ6NmpnSXRUQVVDeWIyRDlpQTk1RGt0Q24xSWNyS0dVZEVjWFNfX0tDRXkyTHRLYXYyZjQwakFWb2Vpd1llRFo2OF9HbHZscXJmUWItM2F0RDd4aEg5WHJXa2EzQUJUNWh6dThqbVZhRDZaemd6R3FNQXpCN3N3aUJ3QXVHazJuTGF1dlNwQmIwRHptM0FPZ2hmcE1IekVKTVNDUWZzVHlTTkRZMDhPQkZvdXlNVDZVZ3VkVk5OS05yVHdiUE9pLUNQTTZjLTBIdi1Hd29oUXozR0VHY3FYMFVqQnU0Q3ZUTTlBV0I3ajR3cTBuU1FVb2thdzNOb2hGLW1XMFVqZEFZM1ZaRmwwLUhHRjZFQUh3dHZDT0RVX3ZtYVNNRXRmOHl4aHhETnVVaDltZDVJeGRwRHk5V1JDcEpEZkVyNTdHX0l1MzZwNDA2VnBrWDVjUDNHSHZLaEI3VHpnSEoxRDJhWkp0YWN6NUxSaUNqMm1VcU42d19hdHVKamNMZ3JoQU1pWDhHY3ZmcldBTzFfUk1DaThFY215S3B1aXJWNDdWczZMcG5paHpPV25NVzdTMzVYRmxQU1FESXhYOHdfWlZwMTdiWVZ0VlQydFZ3a2d1LWNLeDhSaU85RDZwcVUyb2YxZVJUSFNBcl9ISFR3b1Y1VjNfQWJtSXFuLXBfVWpLdGpkUkVSYy1nYkxzenZDR2tYeFAwWVREU0p6dndkcUJzUV9TZXE2X010M1k4dDE5aVVCR3ZsMUVWMXp3QnRFUjhjUzFacjB6a2ZZRWtaeFdlV2dMU3ZrV2F4VnBTd2J1NlJ4dGtwY1hEMktMUVFJVzRmTjBOVEFLY1k4eHczZlYzeTBrUkthWlRmcUt5WklNRjFFZUtSdXJHRzQ3RGJEZEFCeXd2cmZyVGFQeV9tcEtaN2QyZFN0TVEtWC1uRy1QVjlIMWl0MWVqdnB6c1NSd19DLVFib2d5TVRTV2p3dXE2VEtrMnhxN1ZRTXNyMGJuMExYYUxTYlpDSExfU3U0cEdYaFFiT2JTZ1Q3SDNvcGNTMEVXQ2xTYnUweGJBTEQwQUVycVRxZWluUlI0VEhNYmlhMWg1MWVoRllPMDNXZ2pVemhLQ01heV9sb3pZZ2U1NEJGSEpCYWExN3lKMXB4NXRxOXFhd2FFR3V1ci1RRkZTbUFjblo3RWlBM3BLUC14V05acm9vNlZ0QXItWFJSbWdPNlhrOEdPdjlyaHpQVUtxMWNzNVRBS2NuWnFjZE85WC1MS255WU5kUHNTdGRJTHFRdFVGSl93RldFRXMzSnBtRm03QWN5ZDBibWNxMFRXS3ZudEU1U0Y1NGxWT1hYZjl3SGktR2c1QndfVjBPMnpCT0h5aTVyWU1VRzBsZUp6RWYwa3RqRHVkeWdyZEFpbmQ3X1VhUzBJLTJzUXhiVzZOdWxSOVVyWFJ0OFNsTzR4RC1SVll0Yk5DVy0zRXAyYmlZaWVuaFBLNERpZVBQSFlON2dPeDVIYXI3bTdTa010bUNCSmRyLXhOY19QLTRHVF9HY1EyTUtKdG80ZVk1TlFMNVZhZnFiYmVscDVuc3VLeXBCT0lDUjhGSkdaRVhKUEVLcWttZDZTby05bV9iQTQwMWI1ZXZkWWU0TzRPclc3ajZzdjBNVm5HSnZZa19zNERZSHBtQmQ3ZVU2d3dhdUVKcVp0cTdxZGdmMUhreGhMZWtWcEdBN2NkTzhXNVVYV3JRU19ieVF3VVZYQWtMbjV0OXhDTklaY0NQakhpS3RJZ3g5V3VSaXBUejlJblBsbTVucGxFaVU3b29sQXBOWnk3THdYU21XU0VQcE1wSTNvdnNqbmN6VjNRQkJmSUFnZFRpZ2MwbEk5cE9xRzFVWVJGRk5zTjZfVnRVajVLMGxPWWIwejZCWnRidVVOY2FQbmdiclJjM012YVg4XzRQWjRPc3ZFalZNdzkyZE1kaldoVWdrRWwyMEdfSUxaSzB4SThtNDA2a1dVcWdHV01lakw3YVBFQVlUOFVGUXN0N2FxNHBfRnprVTRSTGtkR3M3eEtuSDhKakZxTnA4dXJLZzZJYWJsSnJGRlg5dXVyeDU4NTBHazdUODk0M2w5ckpRd0NyeDVGbG9DSV9uejhCMDJudGV1anE5aXU0eF9hTndKZ1E3RGt0Q255NHdWdVY1YzRVM3Fjc3kyLVZCX05EVFhtNHl4QzRJbWt6bFpXMU11Z1JZcnNtMU5EUkdWOTZSLWM1SlVKMkJaY3ZueEp5bWtHRGxpRUlHVGJyUlJMcjFONUM5RUdWcXhkbnljNWg5V2wyRFNCWTNRYllxblAyeEdCVVNYcW5SVkxxWHlXZjdBSjhLSVpCYzNfMUdMWTN5RVdpS0VTLVVWeHJxNkVNUWFNTXF0NUFqeEZpYU1MTmxSdVdoOHNOWkxmb3g4WGlPVnNJR3dDS3ZYVkpUc1VSYU1fR3lMVWlPNk1pSVhLUjlmTkVBdUxmajc5bnNqYWpTMUUtbmVobVZldzVsbHh5TXNpSXZKZFdYNVh2c3NtYXpnUFBSYm50TVl6eDBVTU9YeDR5dllqZnNweDlhWF9NRmFVcVd5YURYX2lxdEV3NFpFZVRqVWZDcnZKVTJWNlhUNFZrclF3c3FyeWp2T0llS0tkTGZGbjM3RS1CY29YaVM2cnNnMGJDZXVfdmR0eU9aUlZfRTVDdkFhbS1jVW51ZzRnYUhfclBVSzJVNW1fV2o3bHVMcDZqeTlzbDZQdXBLTlZDaUczakU0RmFnN0tydDh3QngzZUJXLTZCNEJJMFRjWTFfRWx6WjhkNGpUam9jQmRhU09FUURWZ0JrWThfZUpsVEdjVkk4YjhzQUV4WkN2UzhxR3p5M0RBczRVX2c4VVVlWnpJZjFYeUFHYkgyeXctNzgzS01PZHVhc0lNbzNWNHkyMF8tekhBV3h2MFFCWUdVUW56d296cC13SFdLZFZYZER4ZlBraU16dzVsb3dPODJGbHkxM1UtRkJxS0hKOUIxRnNvbjY0WnMxazFzaUJUdEI0aDJRQzBsUG1IelVjcmpBTHBxQlRxOUR3MDROSVpNVGNzVVRDS3pzb0hyamNfbmlCT0JMOEhNVldXa2g3TWJSSmZQZm1BVFFhUVFrUjF5QVlpVENnMktxYm5zQTM2LXpULVpxaXM1ZUpVMzR6UlVrMjVEaGxlN3hRXzV1YjQ4TnV3WFNnR3NCTU5PS0Z6T3NWZVZNSGVWYXUwRk1oZ0lCeGFobTBCa29aazdtV05XX1JPSDVocEtjVW1SNmRpd2g5ZWJwc1JWcmZPdFZIUUpiSGo3a2pVSEJZeXM5ZUhOMDFjWW5TT0JFbVlrUXBFQmhmX29EVVY3Yi13elRZR0xQR2syUTVVNFRMSlV0ZlQwaDIwUmMxejJoWjQzOE51OTBva01uUWhsbGxXelRyY1QyaTFaaW1meENEYi1aQXFiMS1YUGRyLTAzUWM3elBweGdqOUJCTXROUmllbnNwTFc0Y3lUTmpOa1FVYk1yaVBJMlFUbWdUOTEydUhyaGJvYkhZeE1RS0dWbVVyNl9VT043S3lHR0txemoteVIwOVh0Vk9PZm5ZenBnZXBRalpLS0pqVDNjZ0htUmNSZC16RnpGTWM1MnRoMkVSZG50VFBmbnpFNzNwRndoRzJBbDJSUzJjVWlaS0QtLXVocnFLWFFVYkYwR1kzdmVWcmV5aHBOT19hX2s0NGd4OTktRGdtOGk1NnZudlE4RTRTbFNudEhhcGhQbk8wVV9SbWdLVmtWeC1JRjR5QUhaOVl3ZnNLWkxSbVlmMFp4UnFHRXBFY1VTbTE0X2tIVmhyNmZuLVZSTXJKQXVRV2NsNlNVWDU5dk53b3JuSVhmRnVGWjRlSTB1NXJ6ejYxbmk2dnNyamY2a1hVNmRqUFk4Nm1manV6ZzVYTWJXLXFxM0tLT25Ua0xfckFtUEttejJIVkdEeU53bl9aMEdqTDRCSkdreVgxemJYV0c5ZXpMRzB5TWVuWGVCVzh6dmRWY1d1dTNHVkp5b1RlVkFpcldyMmJnbVZRSG5GajV6TEhyOS10U0hRdjJRQzNRYUJDQzhHY1ZwVEpvN18yS2hFZGlzY0xSXy04T2tzc19maU81eUhGUmRNMTJBd2hnLm13T2ZEX0FWWGwxU0FuODlsbVByMDNrMFpiS1djS3Q4cHZXem9OSjcyU21RcGh6S3g5OHl6UmVwSi1CUkR3ZlpyZkpuZHJuVWRKaXJ1ell3OGxXakRB"},
		"_eventId":    {"submit"},
		"geolocation": {""},
	})

	data, _ := io.ReadAll(resp.Body)

	invalidCredentials := strings.Contains(string(data), "Authentication attempt has failed, likely due to invalid credentials. Please verify and try again.")
	error_type := strings.Contains(string(data), "<div id=\"loginErrorsPanel\" class=\"banner banner-danger alert alert-danger banner-dismissible\">")

	if invalidCredentials || error_type {
		return true
	} else {
		return false
	}
}