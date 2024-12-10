#!/bin/bash
echo "Deploying application..."

# Stop any existing process
pkill app || true

# Unzip the build artifact
unzip -o app.zip -d /home/ec2-user/

# Start the application
nohup /home/ec2-user/app &

echo "Application deployed successfully!"
