type SensorReading {
    +mint:doc:"Timestamp the reading was taken at"
    +mint:validate:date_in_past
    +mint:transform:date_in_utc
    datetime Timestamp = 0;

    +mint:doc:"Location contains a reference to the specified"
    +mint:doc:"location of this forecast"
    Location Location = 1;

    +mint:doc:"Temperature contained in this reading"
    float32 Temperature = 2;

    +mint:doc:"Humidity contained in this reading"
    float32 Humidity = 3;

    +mint:doc:"PM2_5 of the air in this reading"
    float32 PM2_5 = 4;

    +mint:doc:"Specific sensor used for this reading"
    uuid Sensor = 5;
}
