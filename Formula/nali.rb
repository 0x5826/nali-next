class Nali < Formula
  desc "An offline tool for querying IP geographic information and CDN provider"
  homepage "https://github.com/0x5826/nali-next"
  version "0.8.4"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/0x5826/nali-next/releases/download/v0.8.4/nali-darwin-arm64-v0.8.4.gz"
      sha256 "302d2aba440e766bb633308f24dbdc37f473f7c0aae8f97de3e44fa106b22379"

      def install
        bin.install "nali-darwin-arm64" => "nali"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/0x5826/nali-next/releases/download/v0.8.4/nali-darwin-amd64-v0.8.4.gz"
      sha256 "b7663e22b6be9b506cf60404dca9abb36a6be17b292392196b034a765741818b"

      def install
        bin.install "nali-darwin-amd64" => "nali"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/0x5826/nali-next/releases/download/v0.8.4/nali-linux-armv8-v0.8.4.gz"
      sha256 "eeeb021eaa89d528a8825de9dfa03573e18372fcc377c597f69c8ec91698eb35"

      def install
        bin.install "nali-linux-armv8" => "nali"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/0x5826/nali-next/releases/download/v0.8.4/nali-linux-amd64-v0.8.4.gz"
      sha256 "cb3aaeb230b200490c408bb91b4987acf7aab74d62290ca08c247008a9e61dc3"

      def install
        bin.install "nali-linux-amd64" => "nali"
      end
    end
  end

  test do
    assert_match "1.1.1.1", shell_output("#{bin}/nali 1.1.1.1")
  end
end
