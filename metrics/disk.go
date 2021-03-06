package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
	fs "github.com/dreitier/cloudmon/storage/fs"
)

type Disk struct {
	status                       prometheus.Gauge
	fileCountExpected            *prometheus.GaugeVec
	fileCount                    *prometheus.GaugeVec
	fileAgeThreshold             *prometheus.GaugeVec
	fileYoungCount               *prometheus.GaugeVec
	latestFileCreationExpectedAt *prometheus.GaugeVec
	latestFileCreatedAt          *prometheus.GaugeVec
	latestFileCreationDuration	 *prometheus.GaugeVec
	latestFileBornAt          	 *prometheus.GaugeVec
	latestFileModifiedAt         *prometheus.GaugeVec
	latestFileArchivedAt         *prometheus.GaugeVec
	latestSize                   *prometheus.GaugeVec
}

func NewDisk(diskName string) *Disk {
	presetLabels := map[string]string{"disk": diskName}
	disk := &Disk{
		status: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "status",
			Help:      "Indicates whether there were any problems collecting metrics for this disk. Any value >0 means that errors occurred.",
			ConstLabels: presetLabels,
		}),
		fileCountExpected: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			// TODO BREAKING: Rename that label to files_maximum_count
			Name:      "file_count_aim",
			Help:      "The amount of backup files expected to be present in this group.",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
		}),
		fileCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "file_count",
			Help:      "The amount of backup files present in this group.",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
			"group",
		}),
		fileAgeThreshold: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			// TODO BREAKING: Rename that label to files_maximum_age_in_seconds
			Name:      "file_age_aim_seconds",
			Help:      "The maximum age (in seconds) that any file in this group should reach.",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
		}),
		fileYoungCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "file_young_count",
			Help:      "The amount of backup files in this group that are younger than the maximum age (file_age_aim_seconds).",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
			"group",
		}),
		latestFileCreationExpectedAt: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			// TODO BREAKING: Rename that label to latest_file_creation_expected_at
			Name:      "latest_creation_aim_seconds",
			Help:      "Unix timestamp on which the latest backup in the corresponding file group should have occurred.",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
		}),
		latestFileCreatedAt: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			// TODO BREAKING: Rename that label to latest_file_created_at
			Name:      "latest_creation_seconds",
			Help:      "Unix timestamp on which the latest backup in the corresponding file group was created.",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
			"group",
		}),
		latestFileCreationDuration: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "latest_file_creation_duration",
			Help:      "Describes how long it took to create the backup file in seconds",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
			"group",
		}),
		latestFileBornAt: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "latest_file_born_at",
			Help:      "Unix timestamp on which the latest file has been initially created",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
			"group",
		}),
		latestFileModifiedAt: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "latest_file_modified_at",
			Help:      "Unix timestamp on which the latest file has been modified",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
			"group",
		}),
		latestFileArchivedAt: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "latest_file_archived_at",
			Help:      "Unix timestamp on which the latest file has been archived",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
			"group",
		}),
		latestSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "latest_size_bytes",
			Help:      "Size (in bytes) of the latest backup in the corresponding file group.",
			ConstLabels: presetLabels,
		}, []string{
			"dir",
			"file",
			"group",
		}),
	}
	registry.MustRegister(disk.status)
	registry.MustRegister(disk.fileCountExpected)
	registry.MustRegister(disk.fileCount)
	registry.MustRegister(disk.fileAgeThreshold)
	registry.MustRegister(disk.fileYoungCount)
	registry.MustRegister(disk.latestFileCreationExpectedAt)
	registry.MustRegister(disk.latestFileCreatedAt)
	registry.MustRegister(disk.latestFileCreationDuration)
	registry.MustRegister(disk.latestFileBornAt)
	registry.MustRegister(disk.latestFileModifiedAt)
	registry.MustRegister(disk.latestFileArchivedAt)
	registry.MustRegister(disk.latestSize)
	return disk
}

func (b *Disk) Drop() {
	registry.Unregister(b.status)
	registry.Unregister(b.fileCountExpected)
	registry.Unregister(b.fileCount)
	registry.Unregister(b.fileAgeThreshold)
	registry.Unregister(b.fileYoungCount)
	registry.Unregister(b.latestFileCreationExpectedAt)
	registry.Unregister(b.latestFileCreatedAt)
	registry.Unregister(b.latestFileCreationDuration)
	registry.Unregister(b.latestFileBornAt)
	registry.Unregister(b.latestFileModifiedAt)
	registry.Unregister(b.latestFileArchivedAt)
	registry.Unregister(b.latestSize)
}

func (b *Disk) resetMetrics() {
	b.fileCountExpected.Reset()
	b.fileCount.Reset()
	b.fileAgeThreshold.Reset()
	b.fileYoungCount.Reset()
	b.latestFileCreationExpectedAt.Reset()
	b.latestFileCreatedAt.Reset()
	b.latestFileCreationDuration.Reset()
	b.latestFileBornAt.Reset()
	b.latestFileModifiedAt.Reset()
	b.latestFileArchivedAt.Reset()
	b.latestSize.Reset()
}

func (b *Disk) DefinitionsMissing() {
	b.status.Set(1)
	b.resetMetrics()
}

func (b *Disk) DefinitionsUpdated() {
	b.status.Set(0)
	b.resetMetrics()
}

func (b *Disk) UpdateFileLimits(dir string, file string, count uint64, age time.Duration, ctime time.Time) {
	b.fileCountExpected.WithLabelValues(dir, file).Set(float64(count))
	b.fileAgeThreshold.WithLabelValues(dir, file).Set(age.Seconds())
	b.latestFileCreationExpectedAt.WithLabelValues(dir, file).Set(float64(ctime.Unix()))
}

func (b *Disk) UpdateFileCounts(dir string, file string, group string, present int, young uint64) {
	b.fileCount.WithLabelValues(dir, file, group).Set(float64(present))
	b.fileYoungCount.WithLabelValues(dir, file, group).Set(float64(young))

	if present == 0 {
		labels := make(map[string]string)
		labels["dir"] = dir
		labels["file"] = file
		labels["group"] = group

		b.deleteLatestFileLabels(labels)
	}
}

func (b *Disk) deleteLatestFileLabels(labels map[string]string) {
	b.latestFileCreatedAt.Delete(labels)
	b.latestFileCreationDuration.Delete(labels)
	b.latestFileBornAt.Delete(labels)
	b.latestFileModifiedAt.Delete(labels)
	b.latestFileArchivedAt.Delete(labels)
	b.latestSize.Delete(labels)
}

func (b *Disk) UpdateLatestFile(dir string, file string, group string, fileInfo *fs.FileInfo, time time.Time) {
	b.latestFileCreatedAt.WithLabelValues(dir, file, group).Set(float64(time.Unix()))
	b.latestFileCreationDuration.WithLabelValues(dir, file, group).Set(float64(fileInfo.ModifiedAt.Unix()) - float64(fileInfo.BornAt.Unix()))
	b.latestFileBornAt.WithLabelValues(dir, file, group).Set(float64(fileInfo.BornAt.Unix()))
	b.latestFileModifiedAt.WithLabelValues(dir, file, group).Set(float64(fileInfo.ModifiedAt.Unix()))
	b.latestFileArchivedAt.WithLabelValues(dir, file, group).Set(float64(fileInfo.ArchivedAt.Unix()))
	b.latestSize.WithLabelValues(dir, file, group).Set(float64(fileInfo.Size))
}

func (b *Disk) DropFile(dir string, file string, group string) {
	labels := make(map[string]string)
	labels["dir"] = dir
	labels["file"] = file
	labels["group"] = group

	b.fileCount.Delete(labels)
	b.fileYoungCount.Delete(labels)

	b.deleteLatestFileLabels(labels)
}