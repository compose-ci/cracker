<h2>:mag: Vulnerabilities of <code>hello-wasm-function:demo</code></h2>

<details open="true"><summary>:package: Image Reference</strong> <code>hello-wasm-function:demo</code></summary>
<table>
<tr><td>digest</td><td><code>sha256:f45d12a9acc152df08c99efdccb9bb96fe240f6555c8bc4ef6eedf598aa4f499</code></td><tr><tr><td>vulnerabilities</td><td><img alt="critical: 1" src="https://img.shields.io/badge/critical-1-8b1924"/> <img alt="high: 0" src="https://img.shields.io/badge/high-0-lightgrey"/> <img alt="medium: 0" src="https://img.shields.io/badge/medium-0-lightgrey"/> <img alt="low: 0" src="https://img.shields.io/badge/low-0-lightgrey"/> <!-- unspecified: 0 --></td></tr>
<tr><td>platform</td><td>linux/arm64</td></tr>
<tr><td>size</td><td>4.1 MB</td></tr>
<tr><td>packages</td><td>10</td></tr>
</table>
</details></table>
</details>

<table>
<tr><td valign="top">
<details><summary><img alt="critical: 1" src="https://img.shields.io/badge/C-1-8b1924"/> <img alt="high: 0" src="https://img.shields.io/badge/H-0-lightgrey"/> <img alt="medium: 0" src="https://img.shields.io/badge/M-0-lightgrey"/> <img alt="low: 0" src="https://img.shields.io/badge/L-0-lightgrey"/> <!-- unspecified: 0 --><strong>stdlib</strong> <code>1.24.0</code> (golang)</summary>

<small><code>pkg:golang/stdlib@1.24.0</code></small><br/>
<a href="https://scout.docker.com/v/CVE-2025-22871?s=golang&n=stdlib&t=golang&vr=%3E%3D1.24.0-0%2C%3C1.24.2"><img alt="critical : CVE--2025--22871" src="https://img.shields.io/badge/CVE--2025--22871-lightgrey?label=critical%20&labelColor=8b1924"/></a> 

<table>
<tr><td>Affected range</td><td><code>>=1.24.0-0<br/><1.24.2</code></td></tr>
<tr><td>Fixed version</td><td><code>1.24.2</code></td></tr>
<tr><td>EPSS Score</td><td><code>0.015%</code></td></tr>
<tr><td>EPSS Percentile</td><td><code>2nd percentile</code></td></tr>
</table>

<details><summary>Description</summary>
<blockquote>

The net/http package improperly accepts a bare LF as a line terminator in chunked data chunk-size lines. This can permit request smuggling if a net/http server is used in conjunction with a server that incorrectly accepts a bare LF as part of a chunk-ext.

</blockquote>
</details>
</details></td></tr>
</table>

