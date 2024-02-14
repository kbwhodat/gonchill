class Gonchill < Formula
  desc "Your Go CLI app that allows you to stream via torrents, with quickness"
  homepage "https://github.com/username/gonchill"
  url "https://github.com/kbwhodat/gonchill/releases/download/1.0.4/gonchill-1.0.4.tar.gz"
  sha256 "fdace76053f5ae2990b79a789b6a1caa6cb7ddcaa26371c313b4cb927865a0f0 "
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "#{bin}/gonchill", "main.go"
  end

  test do
    system "#{bin}/gonchill", "--version"
  end
end
