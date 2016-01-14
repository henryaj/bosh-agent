zip bosh-agent-windows.zip service_wrapper.{xml,exe} bosh-agent.exe agent.json settings.json
s3cmd put -P bosh-agent-windows.zip s3://windows-bosh-integration/
