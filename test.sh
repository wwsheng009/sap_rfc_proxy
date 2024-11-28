curl -X POST http://localhost:8080/rfc/call?fname=STFC_STRUCTURE \
-H "Content-Type: application/json" \
-d '{
"IMPORTSTRUCT":{
        "RFCFLOAT": 1.23456789,
        "RFCCHAR1": "A",
        "RFCCHAR2": "BC",
        "RFCCHAR4": "ÄBC",
        "RFCINT1":  1,
        "RFCINT2":  2,
        "RFCINT4":  999999999,
        "RFCTIME":  "080102",
        "RFCDATE":  "20241210",
        "RFCDATA1": "HELLÖ SÄP",
        "RFCDATA2": "DATA222"
    }
}'


curl -X POST http://localhost:8080/rfc/call?fname=/SDF/E2E_DISPATCHED_COLLECTOR \
-H "Content-Type: application/json" \
-d '{
    "DATA_PROVIDER": "/SDF/E2E_ICM_INFO"
}'

curl -X GET http://localhost:8080/rfc/call?fname=STFC_STRUCTURE