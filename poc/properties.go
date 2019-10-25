package poc

import (
	"errors"
	"log"
)

const (
	STRING_LIST_PROPERTY_DELIMITER = ","

	DEFAULT_CHUNK_PART_NONCES           = 960000
	DEFAULT_USE_OPEN_CL                 = true
	DEFAULT_PLATFORM_ID                 = 0
	DEFAULT_DEVICE_ID                   = 0
	DEFAULT_POOL_MINING                 = true
	DEFAULT_FORCE_LOCAL_TARGET_DEADLINE = false
	DEFAULT_DYNAMIC_TARGET_DEADLINE     = false
	DEFAULT_TARGET_DEADLINE             = ^int64(0)
	DEFAULT_SOLO_SERVER                 = "http://localhost:8125"
	DEFAULT_READ_PROGRESS_PER_ROUND     = 9
	DEFAULT_REFRESH_INTERVAL            = 2000
	DEFAULT_CONNECTION_TIMEOUT          = 18000
	DEFAULT_WINNER_RETRIES_ON_ASYNC     = 4
	DEFAULT_WINNER_RETRY_INTERVAL_IN_MS = 4000
	DEFAULT_SCAN_PATHS_EVERY_ROUND      = true
	DEFAULT_BYTE_UNIT_DECIMAL           = true
	DEFAULT_LIST_PLOT_FILES             = false
	DEFAULT_SHOW_DRIVE_INFO             = false
	DEFAULT_SHOW_SKIPPED_DEADLINES      = true
	DEFAULT_READER_THREADS              = 0
	DEFAULT_DEBUG                       = false
	DEFAULT_WRITE_LOG_FILE              = false
	DEFAULT_UPDATE_MINING_INFO          = true
	DEFAULT_LOG_FILE_PATH               = "log/jminer.log.txt"
)

type CoreProperties struct {
	readProgressPerRound     int
	refreshInterval          int64
	connectionTimeout        int64
	winnerRetriesOnAsync     int
	winnerRetryIntervalInMs  int64
	scanPathsEveryRound      bool
	poolMining               bool
	forceLocalTargetDeadline bool
	dynamicTargetDeadline    bool
	targetDeadline           int64
	plotPaths                []string
	chunkPartNonces          int64
	useOpenCl                bool
	deviceId                 int
	platformId               int
	walletServer             string
	numericAccountId         string
	soloServer               string
	passPhrase               string
	poolServer               string
	byteUnitDecimal          bool
	listPlotFiles            bool
	showDriveInfo            bool
	showSkippedDeadlines     bool
	readerThreads            int
	writeLogFile             bool
	debug                    bool
	logFilePath              string
	logPatternFile           string
	logPatternConsole        string
	updateMiningInfo         bool
}

func NewCoreProperties(c *PropertiesConfig) (*CoreProperties, error) {
	m := &CoreProperties{
		readProgressPerRound:     c.ReadProgressPerRound,
		refreshInterval:          c.RefreshInterval,
		connectionTimeout:        c.ConnectionTimeout,
		winnerRetriesOnAsync:     c.WinnerRetriesOnAsync,
		winnerRetryIntervalInMs:  c.WinnerRetryIntervalInMs,
		scanPathsEveryRound:      c.ScanPathsEveryRound,
		poolMining:               c.PoolMining,
		forceLocalTargetDeadline: c.ForceLocalTargetDeadline,
		dynamicTargetDeadline:    c.DynamicTargetDeadline,
		targetDeadline:           c.TargetDeadline,
		plotPaths:                c.PlotPaths,
		chunkPartNonces:          c.ChunkPartNonces,
		useOpenCl:                c.UseOpenCl,
		deviceId:                 c.DeviceId,
		platformId:               c.PlatformId,
		walletServer:             c.WalletServer,
		numericAccountId:         c.NumericAccountId,
		soloServer:               c.SoloServer,
		passPhrase:               c.PassPhrase,
		poolServer:               c.PoolServer,
		byteUnitDecimal:          c.ByteUnitDecimal,
		listPlotFiles:            c.ListPlotFiles,
		showDriveInfo:            c.ShowDriveInfo,
		showSkippedDeadlines:     c.ShowSkippedDeadlines,
		readerThreads:            c.ReaderThreads,
		writeLogFile:             c.WriteLogFile,
		debug:                    c.Debug,
		logFilePath:              c.LogFilePath,
		logPatternFile:           c.LogPatternFile,
		logPatternConsole:        c.LogPatternConsole,
		updateMiningInfo:         c.UpdateMiningInfo,
	}
	if c.ReadProgressPerRound == 0 {
		m.readProgressPerRound = DEFAULT_READ_PROGRESS_PER_ROUND
	}
	if c.RefreshInterval == 0 {
		m.refreshInterval = DEFAULT_REFRESH_INTERVAL
	}
	if c.ConnectionTimeout == 0 {
		m.connectionTimeout = DEFAULT_CONNECTION_TIMEOUT
	}
	if c.WinnerRetriesOnAsync == 0 {
		m.winnerRetriesOnAsync = DEFAULT_WINNER_RETRIES_ON_ASYNC
	}
	if c.WinnerRetryIntervalInMs == 0 {
		m.winnerRetryIntervalInMs = DEFAULT_WINNER_RETRY_INTERVAL_IN_MS
	}
	if !c.ScanPathsEveryRound {
		m.scanPathsEveryRound = DEFAULT_SCAN_PATHS_EVERY_ROUND
	}
	if !c.PoolMining {
		m.poolMining = DEFAULT_POOL_MINING
	}
	if !c.ForceLocalTargetDeadline {
		m.forceLocalTargetDeadline = DEFAULT_FORCE_LOCAL_TARGET_DEADLINE
	}
	if !c.DynamicTargetDeadline {
		m.dynamicTargetDeadline = DEFAULT_DYNAMIC_TARGET_DEADLINE
	}
	if c.TargetDeadline == 0 {
		m.targetDeadline = DEFAULT_TARGET_DEADLINE
	}
	if len(c.PlotPaths) == 0 {
		m.plotPaths = make([]string, 0, 256)
	}
	if c.ChunkPartNonces == 0 {
		m.chunkPartNonces = DEFAULT_CHUNK_PART_NONCES
	}
	if !c.UseOpenCl {
		m.useOpenCl = DEFAULT_USE_OPEN_CL
	}
	if c.DeviceId == 0 {
		m.deviceId = DEFAULT_DEVICE_ID
	}
	if c.PlatformId == 0 {
		m.platformId = DEFAULT_PLATFORM_ID
	}
	if c.WalletServer == "" {
		log.Print("Winner and PoolInfo feature disabled, property 'walletServer' undefined!")
	}
	if c.NumericAccountId == "" {
		return nil, errors.New("property 'numericAccountId' is required for pool-mining!")
	}
	if c.SoloServer == "" {
		m.soloServer = DEFAULT_SOLO_SERVER
	}
	if c.PassPhrase == "" {
		return nil, errors.New("property 'passPhrase' is required for solo-mining!")
	}
	if c.PoolServer == "" {
		return nil, errors.New("property 'poolServer' is required for pool-mining!")
	}
	if !c.ByteUnitDecimal {
		m.byteUnitDecimal = DEFAULT_BYTE_UNIT_DECIMAL
	}
	if !c.ListPlotFiles {
		m.listPlotFiles = DEFAULT_LIST_PLOT_FILES
	}
	if !c.ShowDriveInfo {
		m.showDriveInfo = DEFAULT_SHOW_DRIVE_INFO
	}
	if !c.ShowSkippedDeadlines {
		m.showSkippedDeadlines = DEFAULT_SHOW_SKIPPED_DEADLINES
	}
	if c.ReaderThreads == 0 {
		m.readerThreads = DEFAULT_READER_THREADS
	}
	if !c.WriteLogFile {
		m.writeLogFile = DEFAULT_WRITE_LOG_FILE
	}
	if !c.Debug {
		m.debug = DEFAULT_DEBUG
	}
	if c.LogFilePath == "" {
		m.logFilePath = DEFAULT_LOG_FILE_PATH
	}

	if !c.UpdateMiningInfo {
		m.updateMiningInfo = DEFAULT_UPDATE_MINING_INFO
	}
	return m, nil
}
func (o *CoreProperties) GetReadProgressPerRound() int {
	if o.readProgressPerRound == 0 {
		o.readProgressPerRound = DEFAULT_READ_PROGRESS_PER_ROUND
	}
	return o.readProgressPerRound
}
func (o *CoreProperties) GetRefreshInterval() int64 {
	if o.refreshInterval == 0 {

		o.refreshInterval = DEFAULT_REFRESH_INTERVAL
	}
	return o.refreshInterval
}
