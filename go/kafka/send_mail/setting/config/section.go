package config

// Config structure
type Config struct {
	SendGrid SendGridSetting `mapstructure:"sendgird"`
	Kafka    KafkaSetting    `mapstructure:"kafka"`
	Smtp     SmtpSetting     `mapstructure:"smtp"`
}

// Smtp setting structure
type SmtpSetting struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	Username    string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	FromAddress string `mapstructure:"from_address"`
}

// Send Grid Setting Structure
type SendGridSetting struct {
	APIKey string `mapstructure:"api_key"`
}

// Kafka Setting Structure
type KafkaSetting struct {
	BootstraperSeverMail string `mapstructure:"bootstrap_server_mail"`
}
