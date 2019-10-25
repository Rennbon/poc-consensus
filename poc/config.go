package poc

type PropertiesConfig struct {
	ReadProgressPerRound     int
	RefreshInterval          int64
	ConnectionTimeout        int64
	WinnerRetriesOnAsync     int
	WinnerRetryIntervalInMs  int64
	ScanPathsEveryRound      bool
	PoolMining               bool
	ForceLocalTargetDeadline bool
	DynamicTargetDeadline    bool
	TargetDeadline           int64
	PlotPaths                []string
	ChunkPartNonces          int64
	UseOpenCl                bool
	DeviceId                 int
	PlatformId               int
	WalletServer             string
	NumericAccountId         string
	SoloServer               string
	PassPhrase               string
	PoolServer               string
	ByteUnitDecimal          bool
	ListPlotFiles            bool
	ShowDriveInfo            bool
	ShowSkippedDeadlines     bool
	ReaderThreads            int
	WriteLogFile             bool
	Debug                    bool
	LogFilePath              string
	LogPatternFile           string
	LogPatternConsole        string
	UpdateMiningInfo         bool
}
