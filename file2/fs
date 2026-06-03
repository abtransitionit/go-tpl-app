package file2

import (
	uri "github.com/abtransitionit/go-core/uri3"
)

var (
	pathNotSudo = []string{"tmp", "mx.txt"}
	pathSudo    = []string{"etc", "mx.txt"}

	resourceNotSudo = uri.UriResource{Scheme: "file", Location: pathNotSudo}
	resourceSudo    = uri.UriResource{Scheme: "file", Location: pathSudo}

	hostLocal     = uri.UriHost{Scheme: "local"}
	hostRemoteSsh = uri.UriHost{Scheme: "ssh", Host: "o1u"}

	data = "Hello yoyo from your framework 🚀"
)

// Expose the operations directly as data
var (
	OpeWriteLocal = uri.Operation{
		UriHost:     hostLocal,
		UriResource: resourceNotSudo,
		Action:      "create",
		Options:     map[string]string{"content": data, "mode": "0644"},
	}
	OpeReadLocal = uri.Operation{
		UriHost:     hostLocal,
		UriResource: resourceNotSudo,
		Action:      "read",
		Options:     map[string]string{"content": data, "mode": "0644"},
	}

	OpeWriteLocalSudo = uri.Operation{
		UriHost:     hostLocal,
		UriResource: resourceSudo,
		Action:      "create",
		Options:     map[string]string{"content": data, "mode": "0644", "sudo": "true"},
	}

	OpeReadLocalSudo = uri.Operation{
		UriHost:     hostLocal,
		UriResource: resourceSudo,
		Action:      "read",
		Options:     map[string]string{"content": data, "mode": "0644", "sudo": "true"},
	}

	OpeWriteRemote = uri.Operation{
		UriHost:     hostRemoteSsh,
		UriResource: resourceNotSudo,
		Action:      "create",
		Options:     map[string]string{"content": data, "mode": "0644"},
	}
	OpeReadRemote = uri.Operation{
		UriHost:     hostRemoteSsh,
		UriResource: resourceNotSudo,
		Action:      "read",
		Options:     map[string]string{"content": data, "mode": "0644"},
	}

	OpeWriteRemoteSudo = uri.Operation{
		UriHost:     hostRemoteSsh,
		UriResource: resourceSudo,
		Action:      "create",
		Options:     map[string]string{"content": data, "mode": "0644", "sudo": "true"},
	}
	OpeReadRemoteSudo = uri.Operation{
		UriHost:     hostRemoteSsh,
		UriResource: resourceSudo,
		Action:      "read",
		Options:     map[string]string{"content": data, "mode": "0644", "sudo": "true"},
	}
)
