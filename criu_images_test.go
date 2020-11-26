package criuimages

import (
	"io"
	"testing"

	"github.com/YLonely/criuimages/types"
	"github.com/golang/protobuf/proto"
)

var testCase = struct {
	mainType string
	subType  string
	entries  []types.MntEntry
}{
	mainType: ImageTypeNormal,
	subType:  "mountpoints",
	entries: []types.MntEntry{
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(437),
			RootDev:     proto.Uint32(70),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x200001),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/proc/scsi"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x1),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(436),
			RootDev:     proto.Uint32(69),
			ParentMntId: proto.Uint32(467),
			Flags:       proto.Uint32(0x200001),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/sys/firmware"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x1),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(435),
			RootDev:     proto.Uint32(62),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x1000002),
			Root:        proto.String("/null"),
			Mountpoint:  proto.String("/proc/sched_debug"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String("size=65536k,mode=755"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/dev/null"),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(434),
			RootDev:     proto.Uint32(62),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x1000002),
			Root:        proto.String("/null"),
			Mountpoint:  proto.String("/proc/timer_list"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String("size=65536k,mode=755"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/dev/null"),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(433),
			RootDev:     proto.Uint32(62),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x1000002),
			Root:        proto.String("/null"),
			Mountpoint:  proto.String("/proc/keys"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String("size=65536k,mode=755"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/dev/null"),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(432),
			RootDev:     proto.Uint32(62),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x1000002),
			Root:        proto.String("/null"),
			Mountpoint:  proto.String("/proc/kcore"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String("size=65536k,mode=755"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/dev/null"),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(431),
			RootDev:     proto.Uint32(68),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x200001),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/proc/asound"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x1),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(430),
			RootDev:     proto.Uint32(67),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x200001),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/proc/acpi"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x1),
		},
		{
			Fstype:      proto.Uint32(1),
			MntId:       proto.Uint32(429),
			RootDev:     proto.Uint32(61),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x200001),
			Root:        proto.String("/sysrq-trigger"),
			Mountpoint:  proto.String("/proc/sysrq-trigger"),
			Source:      proto.String("proc"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(1),
			MntId:       proto.Uint32(428),
			RootDev:     proto.Uint32(61),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x200001),
			Root:        proto.String("/sys"),
			Mountpoint:  proto.String("/proc/sys"),
			Source:      proto.String("proc"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(1),
			MntId:       proto.Uint32(427),
			RootDev:     proto.Uint32(61),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x200001),
			Root:        proto.String("/irq"),
			Mountpoint:  proto.String("/proc/irq"),
			Source:      proto.String("proc"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(1),
			MntId:       proto.Uint32(426),
			RootDev:     proto.Uint32(61),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x200001),
			Root:        proto.String("/fs"),
			Mountpoint:  proto.String("/proc/fs"),
			Source:      proto.String("proc"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(1),
			MntId:       proto.Uint32(425),
			RootDev:     proto.Uint32(61),
			ParentMntId: proto.Uint32(462),
			Flags:       proto.Uint32(0x200001),
			Root:        proto.String("/bus"),
			Mountpoint:  proto.String("/proc/bus"),
			Source:      proto.String("proc"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(6),
			MntId:       proto.Uint32(424),
			RootDev:     proto.Uint32(63),
			ParentMntId: proto.Uint32(463),
			Flags:       proto.Uint32(0x20000a),
			Root:        proto.String("/0"),
			Mountpoint:  proto.String("/dev/console"),
			Source:      proto.String("devpts"),
			Options:     proto.String("gid=5,mode=620,ptmxmode=666,newinstance"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(480),
			RootDev:     proto.Uint32(43),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/user.slice/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/devices"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("devices"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(24),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/devices"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(479),
			RootDev:     proto.Uint32(42),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/hugetlb"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("hugetlb"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(23),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/hugetlb"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(478),
			RootDev:     proto.Uint32(41),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/sys/fs/cgroup/rdma"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("rdma"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(22),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/rdma"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(477),
			RootDev:     proto.Uint32(40),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/user.slice/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/memory"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("memory"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(21),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/memory"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(476),
			RootDev:     proto.Uint32(39),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/net_cls,net_prio"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("net_cls,net_prio"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(20),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/net_cls,net_prio"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(475),
			RootDev:     proto.Uint32(38),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/user.slice/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/blkio"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("blkio"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(19),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/blkio"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(474),
			RootDev:     proto.Uint32(37),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/user.slice/user-0.slice/session-1.scope/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/pids"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("pids"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(18),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/pids"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(473),
			RootDev:     proto.Uint32(36),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/perf_event"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("perf_event"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(17),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/perf_event"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(472),
			RootDev:     proto.Uint32(35),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/user.slice/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/cpu,cpuacct"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("cpu,cpuacct"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(16),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/cpu,cpuacct"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(471),
			RootDev:     proto.Uint32(34),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/cpuset"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("cpuset"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(15),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/cpuset"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(470),
			RootDev:     proto.Uint32(33),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/freezer"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("freezer"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(14),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/freezer"),
		},
		{
			Fstype:      proto.Uint32(12),
			MntId:       proto.Uint32(469),
			RootDev:     proto.Uint32(31),
			ParentMntId: proto.Uint32(468),
			Flags:       proto.Uint32(0x28000f),
			Root:        proto.String("/user.slice/user-0.slice/session-1.scope/test"),
			Mountpoint:  proto.String("/sys/fs/cgroup/systemd"),
			Source:      proto.String("cgroup"),
			Options:     proto.String("xattr,name=systemd"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(11),
			SbFlags:     proto.Uint32(0x0),
			ExtKey:      proto.String("/sys/fs/cgroup/systemd"),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(468),
			RootDev:     proto.Uint32(66),
			ParentMntId: proto.Uint32(467),
			Flags:       proto.Uint32(0x20000e),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/sys/fs/cgroup"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String("mode=755"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(2),
			MntId:       proto.Uint32(467),
			RootDev:     proto.Uint32(65),
			ParentMntId: proto.Uint32(461),
			Flags:       proto.Uint32(0x20000f),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/sys"),
			Source:      proto.String("sysfs"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x1),
		},
		{
			Fstype:      proto.Uint32(14),
			MntId:       proto.Uint32(466),
			RootDev:     proto.Uint32(60),
			ParentMntId: proto.Uint32(463),
			Flags:       proto.Uint32(0x20000e),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/dev/mqueue"),
			Source:      proto.String("mqueue"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(465),
			RootDev:     proto.Uint32(64),
			ParentMntId: proto.Uint32(463),
			Flags:       proto.Uint32(0x20000e),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/dev/shm"),
			Source:      proto.String("shm"),
			Options:     proto.String("size=65536k"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(6),
			MntId:       proto.Uint32(464),
			RootDev:     proto.Uint32(63),
			ParentMntId: proto.Uint32(463),
			Flags:       proto.Uint32(0x20000a),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/dev/pts"),
			Source:      proto.String("devpts"),
			Options:     proto.String("gid=5,mode=620,ptmxmode=666,newinstance"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(5),
			MntId:       proto.Uint32(463),
			RootDev:     proto.Uint32(62),
			ParentMntId: proto.Uint32(461),
			Flags:       proto.Uint32(0x1000002),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/dev"),
			Source:      proto.String("tmpfs"),
			Options:     proto.String("size=65536k,mode=755"),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(1),
			MntId:       proto.Uint32(462),
			RootDev:     proto.Uint32(61),
			ParentMntId: proto.Uint32(461),
			Flags:       proto.Uint32(0x200000),
			Root:        proto.String("/"),
			Mountpoint:  proto.String("/proc"),
			Source:      proto.String("proc"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(0),
			SbFlags:     proto.Uint32(0x0),
		},
		{
			Fstype:      proto.Uint32(0),
			MntId:       proto.Uint32(461),
			RootDev:     proto.Uint32(8388610),
			ParentMntId: proto.Uint32(423),
			Flags:       proto.Uint32(0x280000),
			Root:        proto.String("/root/test/java-cs-test/rootfs"),
			Mountpoint:  proto.String("/"),
			Source:      proto.String("/dev/sda2"),
			Options:     proto.String(""),
			SharedId:    proto.Uint32(0),
			MasterId:    proto.Uint32(4294967295),
			SbFlags:     proto.Uint32(0x0),
		},
	},
}

func TestImageRead(t *testing.T) {
	i, err := New("test/mountpoints-13.img")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer i.Close()
	if i.MainType() != ImageTypeNormal || i.SubType() != "mountpoints" {
		t.Error("wrong image type")
		return
	}
	entries := make([]*types.MntEntry, 0, len(testCase.entries))
	for {
		e := &types.MntEntry{}
		err = i.ReadOne(e)
		if err != nil {
			if err != io.EOF {
				t.Error(err.Error())
			}
			break
		}
		entries = append(entries, e)
	}
	if len(entries) != len(testCase.entries) {
		t.Errorf("readed %d lines of data, expected %d lines\n", len(entries), len(testCase.entries))
		return
	}
	t.Logf("lines of entries %d\n", len(entries))
	for i, e := range entries {
		teste := &testCase.entries[i]
		if e.GetFstype() != teste.GetFstype() || e.GetMntId() != teste.GetMntId() || e.GetRootDev() != teste.GetRootDev() ||
			e.GetParentMntId() != teste.GetParentMntId() || e.GetFlags() != teste.GetFlags() || e.GetRoot() != teste.GetRoot() ||
			e.GetMountpoint() != teste.GetMountpoint() || e.GetSource() != teste.GetSource() || e.GetOptions() != teste.GetOptions() ||
			e.GetSharedId() != teste.GetSharedId() || e.GetMasterId() != teste.GetMasterId() || e.GetSbFlags() != teste.GetSbFlags() ||
			e.GetExtKey() != teste.GetExtKey() {
			t.Errorf("mismatch entry at index %d\n", i)
			t.Errorf("\nFstype:%v\nMntId:%v\nRootDev:%v\nParentMntId:%v\nFlags:%v\nRoot:%v\nMountpoint:%v\nSource:%v\nOptions:%v\nSharedId:%v\nMasterId:%v\nSbFlags:%v\nExtKey:%v\n",
				e.GetFstype(), e.GetMntId(), e.GetRootDev(), e.GetParentMntId(), e.GetFlags(), e.GetRoot(), e.GetMountpoint(), e.GetSource(), e.GetOptions(),
				e.GetSharedId(), e.GetMasterId(), e.GetSbFlags(), e.GetExtKey())
			return
		}
	}
}