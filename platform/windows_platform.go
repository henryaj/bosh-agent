package platform

import (
	"encoding/json"
	"path/filepath"

	boshdpresolv "github.com/cloudfoundry/bosh-agent/infrastructure/devicepathresolver"
	boshcert "github.com/cloudfoundry/bosh-agent/platform/cert"
	boshdevutil "github.com/cloudfoundry/bosh-agent/platform/deviceutil"
	boshstats "github.com/cloudfoundry/bosh-agent/platform/stats"
	boshvitals "github.com/cloudfoundry/bosh-agent/platform/vitals"
	boshsettings "github.com/cloudfoundry/bosh-agent/settings"
	boshdir "github.com/cloudfoundry/bosh-agent/settings/directories"
	boshdirs "github.com/cloudfoundry/bosh-agent/settings/directories"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type windowsPlatform struct {
	collector          boshstats.Collector
	fs                 boshsys.FileSystem
	cmdRunner          boshsys.CmdRunner
	compressor         boshcmd.Compressor
	copier             boshcmd.Copier
	dirProvider        boshdirs.Provider
	vitalsService      boshvitals.Service
	devicePathResolver boshdpresolv.DevicePathResolver
	logger             boshlog.Logger
	certManager        boshcert.Manager
	cdutil             boshdevutil.DeviceUtil
}

func NewWindowsPlatform(
	collector boshstats.Collector,
	fs boshsys.FileSystem,
	cmdRunner boshsys.CmdRunner,
	dirProvider boshdirs.Provider,
	devicePathResolver boshdpresolv.DevicePathResolver,
	logger boshlog.Logger,
	cdutil boshdevutil.DeviceUtil,
) Platform {
	return &windowsPlatform{
		fs:                 fs,
		cmdRunner:          cmdRunner,
		collector:          collector,
		compressor:         boshcmd.NewTarballCompressor(cmdRunner, fs),
		copier:             boshcmd.NewCpCopier(cmdRunner, fs, logger),
		dirProvider:        dirProvider,
		devicePathResolver: devicePathResolver,
		vitalsService:      boshvitals.NewService(collector, dirProvider),
		certManager:        boshcert.NewDummyCertManager(fs, cmdRunner, logger),
		cdutil:             cdutil,
	}
}

func (p windowsPlatform) GetFs() (fs boshsys.FileSystem) {
	return p.fs
}

func (p windowsPlatform) GetRunner() (runner boshsys.CmdRunner) {
	return p.cmdRunner
}

func (p windowsPlatform) GetCompressor() (compressor boshcmd.Compressor) {
	return p.compressor
}

func (p windowsPlatform) GetCopier() (copier boshcmd.Copier) {
	return p.copier
}

func (p windowsPlatform) GetDirProvider() (dirProvider boshdir.Provider) {
	return p.dirProvider
}

func (p windowsPlatform) GetVitalsService() (service boshvitals.Service) {
	return p.vitalsService
}

func (p windowsPlatform) GetDevicePathResolver() (devicePathResolver boshdpresolv.DevicePathResolver) {
	return p.devicePathResolver
}

func (p windowsPlatform) SetupRuntimeConfiguration() (err error) {
	return
}

func (p windowsPlatform) CreateUser(username, password, basePath string) (err error) {
	return
}

func (p windowsPlatform) AddUserToGroups(username string, groups []string) (err error) {
	return
}

func (p windowsPlatform) DeleteEphemeralUsersMatching(regex string) (err error) {
	return
}

func (p windowsPlatform) SetupRootDisk(ephemeralDiskPath string) (err error) {
	return
}

func (p windowsPlatform) SetupSSH(publicKey, username string) (err error) {
	return
}

func (p windowsPlatform) SetUserPassword(user, encryptedPwd string) (err error) {
	return
}

func (p windowsPlatform) SetupHostname(hostname string) (err error) {
	return
}

func (p windowsPlatform) SetupNetworking(networks boshsettings.Networks) (err error) {
	return
}

func (p windowsPlatform) GetConfiguredNetworkInterfaces() (interfaces []string, err error) {
	return
}

func (p windowsPlatform) GetCertManager() (certManager boshcert.Manager) {
	return p.certManager
}

func (p windowsPlatform) SetupLogrotate(groupName, basePath, size string) (err error) {
	return
}

func (p windowsPlatform) SetTimeWithNtpServers(servers []string) (err error) {
	return
}

func (p windowsPlatform) SetupEphemeralDiskWithPath(devicePath string) (err error) {
	return
}

func (p windowsPlatform) SetupRawEphemeralDisks(devices []boshsettings.DiskSettings) (err error) {
	return
}

func (p windowsPlatform) SetupDataDir() error {
	return nil
}

func (p windowsPlatform) SetupTmpDir() error {
	return nil
}

func (p windowsPlatform) MountPersistentDisk(diskSettings boshsettings.DiskSettings, mountPoint string) (err error) {
	return
}

func (p windowsPlatform) UnmountPersistentDisk(diskSettings boshsettings.DiskSettings) (didUnmount bool, err error) {
	return
}

func (p windowsPlatform) GetEphemeralDiskPath(diskSettings boshsettings.DiskSettings) string {
	return "/dev/sdb"
}

func (p windowsPlatform) GetFileContentsFromCDROM(diskPath string) ([]byte, error) {
	return p.fs.ReadFile("D:\\env.json")
}

func (p windowsPlatform) GetFilesContentsFromDisk(diskPath string, fileNames []string) (contents [][]byte, err error) {
	return
}

func (p windowsPlatform) MigratePersistentDisk(fromMountPoint, toMountPoint string) (err error) {
	return
}

func (p windowsPlatform) IsMountPoint(path string) (result bool, err error) {
	return
}

func (p windowsPlatform) IsPersistentDiskMounted(diskSettings boshsettings.DiskSettings) (bool, error) {
	return true, nil
}

func (p windowsPlatform) StartMonit() (err error) {
	return
}

func (p windowsPlatform) SetupMonitUser() (err error) {
	return
}

func (p windowsPlatform) GetMonitCredentials() (username, password string, err error) {
	return
}

func (p windowsPlatform) PrepareForNetworkingChange() error {
	return nil
}

func (p windowsPlatform) GetDefaultNetwork() (boshsettings.Network, error) {
	var network boshsettings.Network

	networkPath := filepath.Join(p.dirProvider.BoshDir(), "dummy-default-network-settings.json")
	contents, err := p.fs.ReadFile(networkPath)
	if err != nil {
		return network, nil
	}

	err = json.Unmarshal([]byte(contents), &network)
	if err != nil {
		return network, bosherr.WrapError(err, "Unmarshal json settings")
	}

	return network, nil
}

func (p windowsPlatform) GetHostPublicKey() (string, error) {
	return "dummy-public-key", nil
}
