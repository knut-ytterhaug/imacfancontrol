#!/bin/bash

DIR=/sys/devices/platform/applesmc.768

function fan3Controller() {
	SENSOR=temp30

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
		echo fan3: $(($READING/5000*5000))

		FAN_SPEED=${THRESHOLDS[$(($READING/5000*5000))]}

		echo $FAN_SPEED > /sys/devices/platform/applesmc.768/fan3_min
	done
}


function fan2Controller() {
	SENSOR=temp31

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
		echo $READING 
		echo fan2: $(($READING/5000*5000))

		FAN_SPEED=${THRESHOLDS[$(($READING/5000*5000))]}

		echo $FAN_SPEED > /sys/devices/platform/applesmc.768/fan2_min
	done
}

function fan1Controller() {
	SENSOR=temp21

	THRESHOLDS=(
		[30000]=1000
		[35000]=1250
		[40000]=1500
		[45000]=1750
		[50000]=2000
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
		echo $READING 
		echo fan1: $(($READING/5000*5000))

		FAN_SPEED=${THRESHOLDS[$(($READING/5000*5000))]}

		echo $FAN_SPEED > /sys/devices/platform/applesmc.768/fan1_min
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
