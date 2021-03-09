#!/bin/bash

DIR=/sys/devices/platform/applesmc.768

SENSOR=temp30

THRESHOLDS=(
	[45000]=1000
	[50000]=1250
	[55000]=1500
	[60000]=1750
	[65000]=2000
	[70000]=2500
	[75000]=2750
	[80000]=3000
	[85000]=3250
	[90000]=3500
	[95000]=3750
	[100000]=4000
)

while sleep 5
do
	read READING <$DIR/${SENSOR}_input
	echo $READING 

	FAN_SPEED=${THRESHOLDS[$(($READING/5000*5000))]}

	echo $FAN_SPEED > /sys/devices/platform/applesmc.768/fan3_min
done
