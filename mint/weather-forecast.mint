type WeatherForecast {
    +mint:doc:"ForecastedAt contains the datetime at which this forecast"
    +mint:doc:"was created"
    +mint:validate:date_in_past
    +mint:transform:date_in_utc
    datetime ForecastedAt = 0;

    +mint:doc:"Location contains a reference to the specified"
    +mint:doc:"location of this forecast"
    Location Location = 1;

    +mint:doc:"Temperature is a float containing the forecasted temperature"
    +mint:doc:"temperature"
    float32 Temperature = 2;

    +mint:doc:"Date this forecast is for"
    +mint:transform:date_in_utc
    datetime ForecastedFor = 3;
}
