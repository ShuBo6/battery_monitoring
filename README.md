# 电池状态监测

## 前置依赖
本程序基于upower的电池检测能力，轮询电池状态，定时发送邮件
## 使用场景
自建pve主机，宿主机采用淘汰下来的旧笔记本。电池状态不稳定，需要一直使用电源适配器供电，本程序用来监测电池以及电源的状态。

## 配置环境变量
```shell
TO_EMAILS=xxx@163.com,xxx@qq.com # 必填，否则程序panic
CONF_PATH=/opt/battery_monitoring/conf/config.yaml #可选填，默认为conf/config.yaml
```
## 配置文件conf/config.yaml 模板如下
```yaml
# zap logger configuration
zap:
  level: 'info'
  format: 'console'
  prefix: '[battery_monitoring]'
  director: 'logs'
  link-name: 'latest_log'
  show-line: true
  encode-level: 'LowercaseColorLevelEncoder'
  stacktrace-key: 'disableStacktrace'
  log-in-console: true

system:
  port: "5080"
  monitoringInterval: "300"

email:
  from: "xx@qq.com"
  name: "xxxxx"
  user: "xxxxx"
  host: "smtp.qq.com"
  secret: "xxxxxxxxx"
  port: 587
```

## 最佳实践
将battery_monitoring程序注册成systemd的自定义服务。（pve宿主机就不折腾容器化了，如果是其他发行版可以试一下）
```shell
## 注册自定义服务（注意：这里必须使用绝对路径，并且声明环境变量，参考：https://blog.csdn.net/yanhanhui1/article/details/117196904）
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
## 重载服务
systemctl daemon-reload
## 开机自启
systemctl enable battery_monitoring
## 启动服务
systemctl start battery_monitoring
```



# 直接使用upower （如果对upower不感冒可以不看）
## upower安装

```shell
# 使用 yum
yum install -y upower 
# 使用 apt
apt install -y upower
```

# upower 帮助文档


```shell

root@pve:~# upower -h
Usage:
  upower [OPTION…] UPower tool

Help Options:
  -h, --help           Show help options

Application Options:
  -e, --enumerate      Enumerate objects paths for devices（枚举设备的路径）
  -d, --dump           Dump all parameters for all objects
  -w, --wakeups        Get the wakeup data
  -m, --monitor        Monitor activity from the power daemon
  --monitor-detail     Monitor with detail
  -i, --show-info      Show information about object path（根据设备路径参数展示对应的设备信息）
  -v, --version        Print version of client and daemon
```
# upower的使用示例
```shell
root@pve:~# upower -e
/org/freedesktop/UPower/devices/line_power_AC
/org/freedesktop/UPower/devices/battery_BAT0
/org/freedesktop/UPower/devices/DisplayDevice

root@pve:~# upower -i /org/freedesktop/UPower/devices/line_power_AC
  native-path:          AC
  power supply:         yes
  updated:              Fri 08 Jul 2022 11:08:28 AM CST (1417 seconds ago)
  has history:          no
  has statistics:       no
  line-power
    warning-level:       none
    online:              yes
    icon-name:          'ac-adapter-symbolic'

root@pve:~# upower -i /org/freedesktop/UPower/devices/battery_BAT0
  native-path:          BAT0
  vendor:               SANYO
  model:                01AV417
  serial:               4116
  power supply:         yes
  updated:              Fri 08 Jul 2022 11:30:28 AM CST (103 seconds ago)
  has history:          yes
  has statistics:       yes
  battery
    present:             yes
    rechargeable:        yes
    state:               charging
    warning-level:       none
    energy:              15.51 Wh
    energy-empty:        0 Wh
    energy-full:         34.16 Wh
    energy-full-design:  41.76 Wh
    energy-rate:         17.029 W
    voltage:             15.084 V
    time to full:        1.1 hours
    percentage:          45%
    capacity:            81.8008%
    technology:          lithium-ion
    icon-name:          'battery-good-charging-symbolic'
  History (charge):
    1657251028	45.000	charging
  History (rate):
    1657251028	17.029	charging

```