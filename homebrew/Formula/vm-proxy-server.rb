require "language/go"

class VmProxyServer < Formula
  desc "VM Proxy Server - Allows hypervisors to be controlled from docker containers"
  homepage "https://github.com/blacktop/vm-proxy"
  url "https://github.com/blacktop/vm-proxy.git",
    :revision => "0373576e55f26e55bd33518f9aeca5cdc4988a53"

  head "https://github.com/blacktop/vm-proxy.git", :shallow => false

  depends_on "go" => :build

  go_resource "github.com/docker/machine" do
    url "https://github.com/docker/machine.git",
      :revision => "e3594a517fc3a84a0ba25fa4814e6cbb86444f58"
  end

  go_resource "github.com/gorilla/mux" do
    url "https://github.com/gorilla/mux.git",
      :revision => "e84fac997f7f9015ca0c5a35bf0e7922070c98cb"
  end

  def install
    contents = Dir["{*,.git,.gitignore}"]
    gopath = buildpath/"gopath"
    (gopath/"src/github.com/blacktop/vm-proxy").install contents

    ENV["GOPATH"] = gopath
    ENV.prepend_create_path "PATH", gopath/"bin"

    Language::Go.stage_deps resources, gopath/"src"

    cd gopath/"src/github.com/blacktop/vm-proxy/server" do
      system "go", "get", "-v"
      system "go", "build", "-o", bin/"vm-proxy-server"
      # bin.install "bin/vm-proxy-server"
    end
  end

  plist_options :manual => "vm-proxy-server"

  def plist; <<-EOS.undent
    <?xml version="1.0" encoding="UTF-8"?>
    <!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
    <plist version="1.0">
    <dict>
      <key>Label</key>
      <string>#{plist_name}</string>
      <key>ProgramArguments</key>
      <array>
          <string>#{opt_bin}/vm-proxy-server</string>
      </array>
      <key>WorkingDirectory</key>
      <string>#{HOMEBREW_PREFIX}</string>
      <key>StandardOutPath</key>
      <string>#{var}/log/vm-proxy-server/vm-proxy-server.log</string>
      <key>StandardErrorPath</key>
      <string>#{var}/log/vm-proxy-server/vm-proxy-server.log</string>
      <key>RunAtLoad</key>
      <true/>
      <key>KeepAlive</key>
      <true/>
    </dict>
    </plist>
    EOS
  end

  test do
    system "#{bin}/vm-proxy-server"
  end
end
