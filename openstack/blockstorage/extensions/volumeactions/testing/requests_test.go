package testing

import (
	"testing"

	"github.com/zhuqinghua/gophercloud/openstack/blockstorage/extensions/volumeactions"
	th "github.com/zhuqinghua/gophercloud/testhelper"
	"github.com/zhuqinghua/gophercloud/testhelper/client"
)

func TestAttach(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockAttachResponse(t)

	options := &volumeactions.AttachOpts{
		MountPoint:   "/mnt",
		Mode:         "rw",
		InstanceUUID: "50902f4f-a974-46a0-85e9-7efc5e22dfdd",
	}
	err := volumeactions.Attach(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestBeginDetaching(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockBeginDetachingResponse(t)

	err := volumeactions.BeginDetaching(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestDetach(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDetachResponse(t)

	err := volumeactions.Detach(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", &volumeactions.DetachOpts{}).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUploadImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	MockUploadImageResponse(t)
	options := &volumeactions.UploadImageOpts{
		ContainerFormat: "bare",
		DiskFormat:      "raw",
		ImageName:       "test",
		Force:           true,
	}

	err := volumeactions.UploadImage(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestReserve(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockReserveResponse(t)

	err := volumeactions.Reserve(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUnreserve(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUnreserveResponse(t)

	err := volumeactions.Unreserve(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestInitializeConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockInitializeConnectionResponse(t)

	options := &volumeactions.InitializeConnectionOpts{
		IP:        "127.0.0.1",
		Host:      "stack",
		Initiator: "iqn.1994-05.com.redhat:17cf566367d2",
		Multipath: gophercloud.Disabled,
		Platform:  "x86_64",
		OSType:    "linux2",
	}
	_, err := volumeactions.InitializeConnection(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options).Extract()
	th.AssertNoErr(t, err)
}

func TestTerminateConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockTerminateConnectionResponse(t)

	options := &volumeactions.TerminateConnectionOpts{
		IP:        "127.0.0.1",
		Host:      "stack",
		Initiator: "iqn.1994-05.com.redhat:17cf566367d2",
		Multipath: gophercloud.Enabled,
		Platform:  "x86_64",
		OSType:    "linux2",
	}
	err := volumeactions.TerminateConnection(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestExtendSize(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockExtendSizeResponse(t)

	options := &volumeactions.ExtendSizeOpts{
		NewSize: 3,
	}

	err := volumeactions.ExtendSize(client.ServiceClient(), "cd281d77-8217-4830-be95-9528227c105c", options).ExtractErr()
	th.AssertNoErr(t, err)
}
