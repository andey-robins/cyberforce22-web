#!/bin/bash
cd ..
rm server.zip
zip -r server.zip goserver/
scp server.zip blueteam@10.0.153.79:
ssh blueteam@10.0.153.79