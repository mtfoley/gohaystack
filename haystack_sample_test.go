package gohaystack

var samplePayload = []byte(`
{
    "meta": { "ver": "3.0" },
    "cols": [ { "name": "id" }, { "name": "ahu" }, { "name": "air" }, { "name": "area" }, { "name": "cmd" }, { "name": "cool" },
        { "name": "costPerHour" }, { "name": "cur" }, { "name": "curStatus" }, { "name": "curVal" }, { "name": "damper" }, { "name": "dis" },
        { "name": "disMacro" }, { "name": "discharge" }, { "name": "effective" }, { "name": "elec" }, { "name": "elecCost" },
        { "name": "elecMeterLoad" }, { "name": "energy" }, { "name": "enum" }, { "name": "equip" }, { "name": "equipRef" },
        { "name": "fan" }, { "name": "geoAddr" }, { "name": "geoCity" }, { "name": "geoCoord" }, { "name": "geoCountry" }, { "name": "geoPostalCode" },
        { "name": "geoState" }, { "name": "geoStreet" }, { "name": "heat" }, { "name": "his" }, { "name": "hisEnd" }, { "name": "hisFunc" },
        { "name": "hisInterval" }, { "name": "hisMode" }, { "name": "hisRollupFunc" }, { "name": "hisSize" }, { "name": "hisStart" },
        { "name": "hvac" }, { "name": "kind" }, { "name": "kwSite" }, { "name": "lights" }, { "name": "lightsGroup" }, { "name": "meter" },
        { "name": "metro" }, { "name": "navName" }, { "name": "occupied" }, { "name": "occupiedEnd" }, { "name": "occupiedStart" }, { "name": "outside" },
        { "name": "phone" }, { "name": "point" }, { "name": "power" }, { "name": "pressure" }, { "name": "primaryFunction" }, { "name": "regionRef" },
        { "name": "return" }, { "name": "rooftop" }, { "name": "sensor" }, { "name": "site" }, { "name": "siteMeter" }, { "name": "sitePoint" },
        { "name": "siteRef" }, { "name": "sp" }, { "name": "stage" }, { "name": "store" }, { "name": "storeNum" }, { "name": "summary" },
        { "name": "tariffHis" }, { "name": "temp" }, { "name": "tz" }, { "name": "unit" }, { "name": "weatherRef" }, { "name": "yearBuilt" },
        { "name": "zone" }, { "name": "mod" } ],
    "rows": [
        {
            "id": "r:p:demo:r:23a44701-a89a6c66 Carytown",
            "area": "n:3149 ft\u00b2",
            "dis": "Carytown",
            "geoAddr": "3504 W Cary St, Richmond, VA",
            "geoCity": "Richmond",
            "geoCoord": "c:37.555385,-77.486903",
            "geoCountry": "US",
            "geoPostalCode": "23221",
            "geoState": "VA",
            "geoStreet": "3504 W Cary St",
            "metro": "Richmond",
            "occupiedEnd": "h:20:00:00",
            "occupiedStart": "h:10:00:00",
            "phone": "804.552.2222",
            "primaryFunction": "Retail Store",
            "regionRef": "r:p:demo:r:23a44701-67faf4db Richmond",
            "site": "m:",
            "store": "m:",
            "storeNum": "n:1",
            "tz": "New_York",
            "weatherRef": "r:p:demo:r:23a44701-1af1bca9 Richmond, VA",
            "yearBuilt": "n:1996",
            "mod": "t:2018-12-12T22:24:01.714Z UTC"
        },
        {
            "id": "r:p:demo:r:23a44701-bbc36976 Tariff His",
            "dis": "Tariff His",
            "equipRef": "r:p:demo:r:23a44701-092f16fa Carytown ElecMeter-Main",
            "his": "m:",
            "hisEnd": "t:2019-03-01T00:00:00-05:00 New_York",
            "hisSize": "n:27",
            "hisStart": "t:2017-01-01T00:00:00-05:00 New_York",
            "kind": "Str",
            "point": "m:",
            "siteRef": "r:p:demo:r:23a44701-a89a6c66 Carytown",
            "sp": "m:",
            "tariffHis": "m:",
            "tz": "New_York",
            "mod": "t:2018-12-12T22:24:01.849Z UTC"
        },
        {
            "id": "r:p:demo:r:23a44701-4ea35663 Carytown RTU-1 ZoneTempSp",
            "air": "m:",
            "cur": "m:",
            "curStatus": "ok",
            "curVal": "n:73 \u00b0F",
            "disMacro": "$equipRef $navName",
            "effective": "m:",
            "equipRef": "r:p:demo:r:23a44701-7265b064 Carytown RTU-1",
            "his": "m:",
            "hisEnd": "t:2019-03-12T23:45:00-04:00 New_York",
            "hisMode": "cov",
            "hisSize": "n:76884",
            "hisStart": "t:2017-01-01T00:00:00-05:00 New_York",
            "kind": "Number",
            "navName": "ZoneTempSp",
            "point": "m:",
            "regionRef": "r:p:demo:r:23a44701-67faf4db Richmond",
            "siteRef": "r:p:demo:r:23a44701-a89a6c66 Carytown",
            "sp": "m:",
            "temp": "m:",
            "tz": "New_York",
            "unit": "\u00b0F",
            "zone": "m:",
            "mod": "t:2018-12-12T22:24:01.724Z UTC"
        } 
    ]
}
`)