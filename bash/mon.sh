#!/bin/bash

DIR=/sys/devices/platform/applesmc.768

function fan3Controller() {
	SENSOR=temp30
	FAN=fan3

	THRESHOLDS=(
		[25000]=1000
		[30000]=1000
		[35000]=1000
		[40000]=1000
		[45000]=1250
		[50000]=1500
		[55000]=1750
		[60000]=2000
		[65000]=2250
		[70000]=2500
		[75000]=3000
		[80000]=3500
		[85000]=4000
		[90000]=4000
		[95000]=4000
		[100000]=4000
	)

	while sleep 5
	do
		read READING <$DIR/${SENSOR}_input
		read CURRENT_FAN_SPEED <$DIR/${FAN}_min
		SUGGESTED_FAN_SPEED=${THRESHOLDS[$(($READING/5000*5000))]}

		if [ $SUGGESTED_FAN_SPEED -lt $CURRENT_FAN_SPEED ]
		then
			SUGGESTED_FAN_SPEED=${THRESHOLDS[$((($READING+2500)/5000*5000))]}
		fi
		echo "${FAN}: $READING - $SUGGESTED_FAN_SPEED"
		echo $SUGGESTED_FAN_SPEED > /sys/devices/platform/applesmc.768/${FAN}_min
	done
}


function fan2Controller() {
	SENSOR=temp31
	FAN=fan2

	THRESHOLDS=(
		[20000]=1000
		[25000]=1000
		[30000]=1250
		[35000]=1500
		[40000]=1750
		[45000]=2000
		[50000]=2250
		[55000]=2500
		[60000]=2750
		[65000]=3000
		[70000]=3250
		[75000]=3500
		[80000]=3750
		[85000]=4000
	)

	while sleep 5
	do
		read READING <$DIR/${SENSOR}_input
		read CURRENT_FAN_SPEED <$DIR/${FAN}_min
		SUGGESTED_FAN_SPEED=${THRESHOLDS[$(($READING/5000*5000))]}

		if [ $SUGGESTED_FAN_SPEED -lt $CURRENT_FAN_SPEED ]
		then
			SUGGESTED_FAN_SPEED=${THRESHOLDS[$((($READING+2500)/5000*5000))]}
		fi
		echo "${FAN}: $READING - $SUGGESTED_FAN_SPEED"
		echo $SUGGESTED_FAN_SPEED > /sys/devices/platform/applesmc.768/${FAN}_min
	done
}

function fan1Controller() {
	SENSOR=temp21
	FAN=fan1

	THRESHOLDS=(
		[30000]=1000
		[35000]=1000
		[40000]=1250
		[45000]=1500
		[50000]=1750
		[55000]=2000
		[60000]=2500
		[65000]=3000
		[70000]=3250
		[75000]=3500
		[80000]=3750
		[85000]=4000
	)

	while sleep 5
	do
		read READING <$DIR/${SENSOR}_input
		read CURRENT_FAN_SPEED <$DIR/${FAN}_min
		SUGGESTED_FAN_SPEED=${THRESHOLDS[$(($READING/5000*5000))]}

		if [ $SUGGESTED_FAN_SPEED -lt $CURRENT_FAN_SPEED ]
		then
			SUGGESTED_FAN_SPEED=${THRESHOLDS[$((($READING+2500)/5000*5000))]}
		fi
		echo "${FAN}: $READING - $SUGGESTED_FAN_SPEED"
		echo $SUGGESTED_FAN_SPEED > /sys/devices/platform/applesmc.768/${FAN}_min
	done
}

fan1Controller&
fan2Controller&
fan3Controller&

while sleep 1
do
	pids="$(jobs -p)"
	if [ "$pids" == "" ]
	then
		break
	fi
done	
