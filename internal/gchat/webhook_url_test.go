package gchat

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewWebhookURL(t *testing.T) {
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name    string
		args    args
		want    *WebhookURL
		wantErr bool
	}{
		{
			name: "valid URL",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/spaceID/messages/",
					RawQuery: "key=key&token=token",
				},
			},
			want: &WebhookURL{
				URL: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/spaceID/messages/",
					RawQuery: "key=key&token=token",
				},
				SpaceID:    "spaceID",
				APIVersion: "v1",
				Key:        "key",
				Token:      "token",
			},
			wantErr: false,
		},
		{
			name: "invalid URL, unsupported API Version",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/badpathstring/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, unsupported API Version 2",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/b1/spaces/spaceID/messages/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, unsupported spaces",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/thisisnotspaces/spaceID/messages/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, SpaceID is empty",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/messages/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid URL, messages string is not in the url path",
			args: args{
				u: &url.URL{
					Scheme:   "https",
					Host:     "chat.googleapis.com",
					Path:     "/v1/spaces/spaceID/",
					RawQuery: "key=key&token=token",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWebhookURL(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWebhookURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWebhookURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
