package file2

import (
	uri "github.com/abtransitionit/go-core/uri2"
)

var (
	pathNotSudo = []string{"tmp", "mx.txt"}
	pathSudo    = []string{"etc", "mx.txt"}

	resourceNotSudo = uri.Resource{Scheme: "file", Path: pathNotSudo}
	resourceSudo    = uri.Resource{Scheme: "file", Path: pathSudo}

	targetLocal     = uri.Host{Scheme: "local"}
	targetRemoteSsh = uri.Host{Scheme: "ssh", Host: "o1u"}

	data = "Hello yoyo from your framework 🚀"
)

// Expose the operations directly as data
var (
	OpeWriteLocal = uri.Operation{
		Host:     targetLocal,
		Resource: resourceNotSudo,
		Action:   "create",
		Options:  map[string]string{"content": data, "mode": "0644"},
	}
	OpeReadLocal = uri.Operation{
		Host:     targetLocal,
		Resource: resourceNotSudo,
		Action:   "read",
		Options:  map[string]string{"content": data, "mode": "0644"},
	}

	OpeWriteLocalSudo = uri.Operation{
		Host:     targetLocal,
		Resource: resourceSudo,
		Action:   "create",
		Options:  map[string]string{"content": data, "mode": "0644", "sudo": "true"},
	}

	OpeReadLocalSudo = uri.Operation{
		Host:     targetLocal,
		Resource: resourceSudo,
		Action:   "read",
		Options:  map[string]string{"content": data, "mode": "0644", "sudo": "true"},
	}

	OpeWriteRemote = uri.Operation{
		Host:     targetRemoteSsh,
		Resource: resourceNotSudo,
		Action:   "create",
		Options:  map[string]string{"content": data, "mode": "0644"},
	}
	OpeReadRemote = uri.Operation{
		Host:     targetRemoteSsh,
		Resource: resourceNotSudo,
		Action:   "read",
		Options:  map[string]string{"content": data, "mode": "0644"},
	}

	OpeWriteRemoteSudo = uri.Operation{
		Host:     targetRemoteSsh,
		Resource: resourceSudo,
		Action:   "create",
		Options:  map[string]string{"content": data, "mode": "0644", "sudo": "true"},
	}
	OpeReadRemoteSudo = uri.Operation{
		Host:     targetRemoteSsh,
		Resource: resourceSudo,
		Action:   "read",
		Options:  map[string]string{"content": data, "mode": "0644", "sudo": "true"},
	}
)
