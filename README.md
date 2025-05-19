<h1 align='center'>DeObFU v1.0</h1>

<h2 align='center'><a href='#about'>About</a> ℹ️ • <a href='#dd'>Description</a> 🔬 • <a href='#ii'>Installation</a> 🛠️ • <a href='#htu'>How To Use</a> 🤚</h2>

<h2 id='about' align="center">About</h2>
<strong>DeObFU</strong> <i>(Deobfuscator & Fingerprint Utility)</i> is a fast and modular CLI tool for decoding, hash identification, and breaking simple ciphers.
<p></p>
<p>It automatically detects common encodings (Base64, Hex, JWT, etc.), identifies hash types (MD5, SHA-1, NTLM…), and attempts to decipher Caesar/ROT-based ciphers.</p>

Whether you're working on penetration tests, bug bounty hunting, or CTF challenges, DeObFU helps you quickly unravel obfuscated strings and understand what you're looking at.

<h2 id='dd' align="center">Description</h2>
<h3><ins>Supported Features:</ins></h3>

🔐 <strong>Encodings:</strong>
<ul>
  <li>Base64 / Base64URL</li>
  <li>Base32</li>
  <li>Base85</li>
  <li>Hex</li>
  <li>JWT (with header/payload decoding)</li>
</ul>

🧬 <strong>Hash Detection:</strong>
<ul>
  <li>MD5</li>  
  <li>SHA-1</li>  
  <li>SHA-256</li>  
  <li>SHA-512</li>  
  <li>NTLM</li>
</ul>

For each hash match, DeObFU also provides ready-to-use john, hashcat commands and resources for cracking.

🧠 <strong>Cipher Decoding:</strong>
<ul>
  <li>Caesar cipher (with both positive and negative shifts)</li>
  <li>ROT13 / ROT47</li>
  <li>Atbash</li>
</ul>

🌀 <strong>Recursive Decoding:</strong>
Automatically decodes nested/stacked encodings (e.g., Base64 → Base32 → Base64 → text)

<h3>🛠 <ins>Flags:</ins></h3>

| Flag                     | Alias    | Description                                              |
|--------------------------|----------|----------------------------------------------------------|
| `--decode`               | `-dcd`   | Try to decode known encodings like Base64, Hex, JWT     |
| `--hash-identify`        | `-hi`    | Identify possible hash type (MD5, SHA-1, NTLM…)         |
| `--decypher`             | `-dcyph` | Try to break simple ciphers like Caesar, ROT, Atbash    |
| `--auto`                 | —        | Automatically run decode and hash-identify              |
| `--recurse-decode`       | `-rdc`   | Recursively decode stacked encodings                    |
| `--string <value>`       | `-s`     | Input string to analyze                                 |
| `--help`                 | `-h`     | Show this help message                                   |

<h2 id='ii' align='center'>Installation Instructions</h2>
<ul>
  <li>📦 <strong><ins>From source (requires Go)</ins></strong></li>
    <pre><code>go install github.com/Kode-n-Rolla/deobfu/cmd/deobfu@latest</code></pre>
    Make sure <code>$GOPATH/bin</code> is in your <code>$PATH</code>.
  <p></p>
  <li><strong>📥 <ins>From binary release (no Go required)</ins></strong></li>
    <ol>
      <li>Go to the <a href=''>Releases</a> page
      <li>Download the appropriate binary for your system:
        <ul>
          <li><a href=''>deobfu.exe</a> for Windows</li>
          <li><a href=''>deobfu</a> for Linux/macOS</li>
        </ul>
      <li>Make it executable:</li>
        <pre><code>chmod +x deobfu
./deobfu --help</code></pre>
    </ol>
</ul>

<h2 id='htu' align='center'>How To Use</h2>
<strong>DeObFU</strong> can be launched in different ways depending on your workflow:
<ul>
  <li>🧠 <strong>Interactive Mode (no arguments)</strong></li>
    Running the tool with no arguments will launch an intuitive CLI prompt.
    <p>You’ll be asked to enter an obfuscated string, and then select the action you want to perform (decode, identify hash, break cipher, etc).</p>
    <pre><code>deobfu</code></pre>
    This is the fastest and most convenient way to use DeObFU casually or during manual analysis.
  <p></p>
  <li>🔍 <strong>With <code>--string</code> argument</strong></li>
    You can provide a string directly:
    <pre><code>deobfu --string aGVsbG8=</code></pre>
    In this case, the tool will prompt you to select the deobfuscation mode interactively.
  <p></p>
  <li>🎯 <strong>With a specific flag + interactive string input</strong></li>
    You can specify the desired operation mode (e.g., decoding, hash detection, etc), and the tool will ask you to enter the string:
    <pre><code>deobfu --decode</code></pre>
  <li>🧪 <strong>Full automation</strong></li>
    If you want to skip interaction completely and let DeObFU run its full logic:
    <pre><code>deobfu --auto --string MFDVM43CI44D2===</code></pre>
    The tool will automatically analyze the string using all relevant modules and output results.
</ul>
