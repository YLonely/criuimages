package criuimages

const (
	/*
	 * Basic multi-file images
	 */

	CRTOOLS_IMAGES_V1 uint32 = 1
	/*
	 * v1.1 has common magic in the head of each image file,
	 * except for inventory
	 */
	CRTOOLS_IMAGES_V1_1 uint32 = 2

	/*
	 * Raw images are images in which data is stored in some
	 * non-crtool format (ip tool dumps, tarballs, etc.)
	 */

	RAW_IMAGE_MAGIC uint32 = 0x0

	/*
	 * Images have the IMG_COMMON_MAGIC in the head. Service files
	 * such as stats and irmap-cache have the IMG_SERVICE_MAGIC.
	 */

	IMG_COMMON_MAGIC  uint32 = 0x54564319 /* Sarov (a.k.a. Arzamas-16) */
	IMG_SERVICE_MAGIC uint32 = 0x55105940 /* Zlatoust */

	/*
	 * The magic-s below correspond to coordinates
	 * of various Russian towns in the NNNNEEEE form.
	 */

	INVENTORY_MAGIC      uint32 = 0x58313116 /* Veliky Novgorod */
	PSTREE_MAGIC         uint32 = 0x50273030 /* Kyiv */
	FDINFO_MAGIC         uint32 = 0x56213732 /* Dmitrov */
	PAGEMAP_MAGIC        uint32 = 0x56084025 /* Vladimir */
	SHMEM_PAGEMAP_MAGIC  uint32 = PAGEMAP_MAGIC
	PAGES_MAGIC          uint32 = RAW_IMAGE_MAGIC
	CORE_MAGIC           uint32 = 0x55053847 /* Kolomna */
	IDS_MAGIC            uint32 = 0x54432030 /* Konigsberg */
	VMAS_MAGIC           uint32 = 0x54123737 /* Tula */
	PIPES_MAGIC          uint32 = 0x56513555 /* Tver */
	PIPES_DATA_MAGIC     uint32 = 0x56453709 /* Dubna */
	FIFO_MAGIC           uint32 = 0x58364939 /* Kirov */
	FIFO_DATA_MAGIC      uint32 = 0x59333054 /* Tosno */
	SIGACT_MAGIC         uint32 = 0x55344201 /* Murom */
	UNIXSK_MAGIC         uint32 = 0x54373943 /* Ryazan */
	INETSK_MAGIC         uint32 = 0x56443851 /* Pereslavl */
	PACKETSK_MAGIC       uint32 = 0x60454618 /* Veliky Ustyug */
	ITIMERS_MAGIC        uint32 = 0x57464056 /* Kostroma */
	POSIX_TIMERS_MAGIC   uint32 = 0x52603957 /* Lipetsk */
	SK_QUEUES_MAGIC      uint32 = 0x56264026 /* Suzdal */
	UTSNS_MAGIC          uint32 = 0x54473203 /* Smolensk */
	CREDS_MAGIC          uint32 = 0x54023547 /* Kozelsk */
	IPC_VAR_MAGIC        uint32 = 0x53115007 /* Samara */
	IPCNS_SHM_MAGIC      uint32 = 0x46283044 /* Odessa */
	IPCNS_MSG_MAGIC      uint32 = 0x55453737 /* Moscow */
	IPCNS_SEM_MAGIC      uint32 = 0x59573019 /* St. Petersburg */
	REG_FILES_MAGIC      uint32 = 0x50363636 /* Belgorod */
	EXT_FILES_MAGIC      uint32 = 0x59255641 /* Usolye */
	FS_MAGIC             uint32 = 0x51403912 /* Voronezh */
	MM_MAGIC             uint32 = 0x57492820 /* Pskov */
	REMAP_FPATH_MAGIC    uint32 = 0x59133954 /* Vologda */
	GHOST_FILE_MAGIC     uint32 = 0x52583605 /* Oryol */
	TCP_STREAM_MAGIC     uint32 = 0x51465506 /* Orenburg */
	EVENTFD_FILE_MAGIC   uint32 = 0x44523722 /* Anapa */
	EVENTPOLL_FILE_MAGIC uint32 = 0x45023858 /* Krasnodar */
	EVENTPOLL_TFD_MAGIC  uint32 = 0x44433746 /* Novorossiysk */
	SIGNALFD_MAGIC       uint32 = 0x57323820 /* Uglich */
	INOTIFY_FILE_MAGIC   uint32 = 0x48424431 /* Volgograd */
	INOTIFY_WD_MAGIC     uint32 = 0x54562009 /* Svetlogorsk (Rauschen) */
	MNTS_MAGIC           uint32 = 0x55563928 /* Petushki */
	NETDEV_MAGIC         uint32 = 0x57373951 /* Yaroslavl */
	NETNS_MAGIC          uint32 = 0x55933752 /* Dolgoprudny */
	TTY_FILES_MAGIC      uint32 = 0x59433025 /* Pushkin */
	TTY_INFO_MAGIC       uint32 = 0x59453036 /* Kolpino */
	TTY_DATA_MAGIC       uint32 = 0x59413026 /* Pavlovsk */
	FILE_LOCKS_MAGIC     uint32 = 0x54323616 /* Kaluga */
	RLIMIT_MAGIC         uint32 = 0x57113925 /* Rostov */
	FANOTIFY_FILE_MAGIC  uint32 = 0x55096122 /* Chelyabinsk */
	FANOTIFY_MARK_MAGIC  uint32 = 0x56506035 /* Yekaterinburg */
	SIGNAL_MAGIC         uint32 = 0x59255647 /* Berezniki */
	PSIGNAL_MAGIC        uint32 = SIGNAL_MAGIC
	NETLINK_SK_MAGIC     uint32 = 0x58005614 /* Perm */
	NS_FILES_MAGIC       uint32 = 0x61394011 /* Nyandoma */
	TUNFILE_MAGIC        uint32 = 0x57143751 /* Kalyazin */
	CGROUP_MAGIC         uint32 = 0x59383330 /* Tikhvin */
	TIMERFD_MAGIC        uint32 = 0x50493712 /* Korocha */
	CPUINFO_MAGIC        uint32 = 0x61404013 /* Nyandoma */
	USERNS_MAGIC         uint32 = 0x55474906 /* Kazan */
	SECCOMP_MAGIC        uint32 = 0x64413049 /* Kostomuksha */
	BINFMT_MISC_MAGIC    uint32 = 0x67343323 /* Apatity */
	AUTOFS_MAGIC         uint32 = 0x49353943 /* Sochi */
	FILES_MAGIC          uint32 = 0x56303138 /* Toropets */
	MEMFD_INODE_MAGIC    uint32 = 0x48453499 /* Dnipro */
	TIMENS_MAGIC         uint32 = 0x43114433 /* Beslan */
	PIDNS_MAGIC          uint32 = 0x61157326 /* Surgut */
	BPFMAP_FILE_MAGIC    uint32 = 0x57506142 /* Alapayevsk */
	BPFMAP_DATA_MAGIC    uint32 = 0x64324033 /* Arkhangelsk */

	IFADDR_MAGIC    uint32 = RAW_IMAGE_MAGIC
	ROUTE_MAGIC     uint32 = RAW_IMAGE_MAGIC
	ROUTE6_MAGIC    uint32 = RAW_IMAGE_MAGIC
	RULE_MAGIC      uint32 = RAW_IMAGE_MAGIC
	TMPFS_IMG_MAGIC uint32 = RAW_IMAGE_MAGIC
	TMPFS_DEV_MAGIC uint32 = RAW_IMAGE_MAGIC
	IPTABLES_MAGIC  uint32 = RAW_IMAGE_MAGIC
	IP6TABLES_MAGIC uint32 = RAW_IMAGE_MAGIC
	NFTABLES_MAGIC  uint32 = RAW_IMAGE_MAGIC
	NETNF_CT_MAGIC  uint32 = RAW_IMAGE_MAGIC
	NETNF_EXP_MAGIC uint32 = RAW_IMAGE_MAGIC

	PAGES_OLD_MAGIC       uint32 = PAGEMAP_MAGIC
	SHM_PAGES_OLD_MAGIC   uint32 = PAGEMAP_MAGIC
	BINFMT_MISC_OLD_MAGIC uint32 = BINFMT_MISC_MAGIC

	/*
	 * These are special files, not exactly images
	 */
	STATS_MAGIC       uint32 = 0x57093306 /* Ostashkov */
	IRMAP_CACHE_MAGIC uint32 = 0x57004059 /* Ivanovo */

	/*
	 * Main magic for kerndat_s structure.
	 */

	KDAT_MAGIC uint32 = 0x57023458 /* Torzhok */
)

var imageTypes = map[uint32]string{
	INVENTORY_MAGIC:      "inventory",
	FDINFO_MAGIC:         "fdinfo",
	REG_FILES_MAGIC:      "reg-files",
	EVENTFD_FILE_MAGIC:   "eventfd",
	EVENTPOLL_FILE_MAGIC: "eventpoll",
	EVENTPOLL_TFD_MAGIC:  "eventpoll-tfd",
	INOTIFY_FILE_MAGIC:   "inotify",
	INOTIFY_WD_MAGIC:     "inotify-wd",
	SIGNALFD_MAGIC:       "signalfd",
	CORE_MAGIC:           "core",
	MM_MAGIC:             "mm",
	PIPES_MAGIC:          "pipes",
	PIPES_DATA_MAGIC:     "pipes-data",
	FIFO_MAGIC:           "fifo",
	FIFO_DATA_MAGIC:      "fifo-data",
	PSTREE_MAGIC:         "pstree",
	IDS_MAGIC:            "ids",
	SIGACT_MAGIC:         "sigacts",
	UNIXSK_MAGIC:         "unixsk",
	INETSK_MAGIC:         "inetsk",
	SK_QUEUES_MAGIC:      "sk-queues",
	ITIMERS_MAGIC:        "itimers",
	CREDS_MAGIC:          "creds",
	FS_MAGIC:             "fs",
	REMAP_FPATH_MAGIC:    "remap-fpath",
	GHOST_FILE_MAGIC:     "ghost-file",
	TCP_STREAM_MAGIC:     "tcp-stream",
	MNTS_MAGIC:           "mountpoints",
	UTSNS_MAGIC:          "utsns",
	TTY_FILES_MAGIC:      "tty",
	TTY_INFO_MAGIC:       "tty-info",
	PACKETSK_MAGIC:       "packetsk",
	NETDEV_MAGIC:         "netdev",
}
