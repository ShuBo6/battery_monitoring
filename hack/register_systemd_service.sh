cat > /usr/lib/systemd/system/battery_monitoring.service << EOF
[Unit]
Description=BatteryMonitoring
After=network.target remote-fs.target nss-lookup.target

[Service]
Type=simple
Environment="TO_EMAILS=xxx@qq.com,xxx@163.com"
Environment="CONF_PATH=/opt/battery_monitoring/conf/config.yaml"
ExecStart=/opt/battery_monitoring/battery_monitoring
ExecStop=kill -9 $(pidof battery_monitoring)
ExecReload=kill -9 $(pidof battery_monitoring) && /opt/battery_monitoring/battery_monitoring

[Install]
WantedBy=multi-user.target
EOF
