package Crypto

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The Crypto class provides access to advanced cryptographic functionalities.
Currently, this includes asymmetric key encryption/decryption, signing/verification, and generating cryptographically secure random bytes, RSA keys, HMAC digests, and self-signed [X509Certificate]s.
[codeblocks]
[gdscript]
var crypto = Crypto.new()

# Generate new RSA key.
var key = crypto.generate_rsa(4096)

# Generate new self-signed certificate with the given key.
var cert = crypto.generate_self_signed_certificate(key, "CN=mydomain.com,O=My Game Company,C=IT")

# Save key and certificate in the user folder.
key.save("user://generated.key")
cert.save("user://generated.crt")

# Encryption
var data = "Some data"
var encrypted = crypto.encrypt(key, data.to_utf8_buffer())

# Decryption
var decrypted = crypto.decrypt(key, encrypted)

# Signing
var signature = crypto.sign(HashingContext.HASH_SHA256, data.sha256_buffer(), key)

# Verifying
var verified = crypto.verify(HashingContext.HASH_SHA256, data.sha256_buffer(), signature, key)

# Checks
assert(verified)
assert(data.to_utf8_buffer() == decrypted)
[/gdscript]
[csharp]
using Godot;
using System.Diagnostics;

Crypto crypto = new Crypto();

// Generate new RSA key.
CryptoKey key = crypto.GenerateRsa(4096);

// Generate new self-signed certificate with the given key.
X509Certificate cert = crypto.GenerateSelfSignedCertificate(key, "CN=mydomain.com,O=My Game Company,C=IT");

// Save key and certificate in the user folder.
key.Save("user://generated.key");
cert.Save("user://generated.crt");

// Encryption
string data = "Some data";
byte[] encrypted = crypto.Encrypt(key, data.ToUtf8Buffer());

// Decryption
byte[] decrypted = crypto.Decrypt(key, encrypted);

// Signing
byte[] signature = crypto.Sign(HashingContext.HashType.Sha256, Data.Sha256Buffer(), key);

// Verifying
bool verified = crypto.Verify(HashingContext.HashType.Sha256, Data.Sha256Buffer(), signature, key);

// Checks
Debug.Assert(verified);
Debug.Assert(data.ToUtf8Buffer() == decrypted);
[/csharp]
[/codeblocks]

*/
type Go [1]classdb.Crypto

/*
Generates a [PackedByteArray] of cryptographically secure random bytes with given [param size].
*/
func (self Go) GenerateRandomBytes(size int) []byte {
	return []byte(class(self).GenerateRandomBytes(gd.Int(size)).Bytes())
}

/*
Generates an RSA [CryptoKey] that can be used for creating self-signed certificates and passed to [method StreamPeerTLS.accept_stream].
*/
func (self Go) GenerateRsa(size int) gdclass.CryptoKey {
	return gdclass.CryptoKey(class(self).GenerateRsa(gd.Int(size)))
}

/*
Generates a self-signed [X509Certificate] from the given [CryptoKey] and [param issuer_name]. The certificate validity will be defined by [param not_before] and [param not_after] (first valid date and last valid date). The [param issuer_name] must contain at least "CN=" (common name, i.e. the domain name), "O=" (organization, i.e. your company name), "C=" (country, i.e. 2 lettered ISO-3166 code of the country the organization is based in).
A small example to generate an RSA key and an X509 self-signed certificate.
[codeblocks]
[gdscript]
var crypto = Crypto.new()
# Generate 4096 bits RSA key.
var key = crypto.generate_rsa(4096)
# Generate self-signed certificate using the given key.
var cert = crypto.generate_self_signed_certificate(key, "CN=example.com,O=A Game Company,C=IT")
[/gdscript]
[csharp]
var crypto = new Crypto();
// Generate 4096 bits RSA key.
CryptoKey key = crypto.GenerateRsa(4096);
// Generate self-signed certificate using the given key.
X509Certificate cert = crypto.GenerateSelfSignedCertificate(key, "CN=mydomain.com,O=My Game Company,C=IT");
[/csharp]
[/codeblocks]
*/
func (self Go) GenerateSelfSignedCertificate(key gdclass.CryptoKey) gdclass.X509Certificate {
	return gdclass.X509Certificate(class(self).GenerateSelfSignedCertificate(key, gd.NewString("CN=myserver,O=myorganisation,C=IT"), gd.NewString("20140101000000"), gd.NewString("20340101000000")))
}

/*
Sign a given [param hash] of type [param hash_type] with the provided private [param key].
*/
func (self Go) Sign(hash_type classdb.HashingContextHashType, hash []byte, key gdclass.CryptoKey) []byte {
	return []byte(class(self).Sign(hash_type, gd.NewPackedByteSlice(hash), key).Bytes())
}

/*
Verify that a given [param signature] for [param hash] of type [param hash_type] against the provided public [param key].
*/
func (self Go) Verify(hash_type classdb.HashingContextHashType, hash []byte, signature []byte, key gdclass.CryptoKey) bool {
	return bool(class(self).Verify(hash_type, gd.NewPackedByteSlice(hash), gd.NewPackedByteSlice(signature), key))
}

/*
Encrypt the given [param plaintext] with the provided public [param key].
[b]Note:[/b] The maximum size of accepted plaintext is limited by the key size.
*/
func (self Go) Encrypt(key gdclass.CryptoKey, plaintext []byte) []byte {
	return []byte(class(self).Encrypt(key, gd.NewPackedByteSlice(plaintext)).Bytes())
}

/*
Decrypt the given [param ciphertext] with the provided private [param key].
[b]Note:[/b] The maximum size of accepted ciphertext is limited by the key size.
*/
func (self Go) Decrypt(key gdclass.CryptoKey, ciphertext []byte) []byte {
	return []byte(class(self).Decrypt(key, gd.NewPackedByteSlice(ciphertext)).Bytes())
}

/*
Generates an [url=https://en.wikipedia.org/wiki/HMAC]HMAC[/url] digest of [param msg] using [param key]. The [param hash_type] parameter is the hashing algorithm that is used for the inner and outer hashes.
Currently, only [constant HashingContext.HASH_SHA256] and [constant HashingContext.HASH_SHA1] are supported.
*/
func (self Go) HmacDigest(hash_type classdb.HashingContextHashType, key []byte, msg []byte) []byte {
	return []byte(class(self).HmacDigest(hash_type, gd.NewPackedByteSlice(key), gd.NewPackedByteSlice(msg)).Bytes())
}

/*
Compares two [PackedByteArray]s for equality without leaking timing information in order to prevent timing attacks.
See [url=https://paragonie.com/blog/2015/11/preventing-timing-attacks-on-string-comparison-with-double-hmac-strategy]this blog post[/url] for more information.
*/
func (self Go) ConstantTimeCompare(trusted []byte, received []byte) bool {
	return bool(class(self).ConstantTimeCompare(gd.NewPackedByteSlice(trusted), gd.NewPackedByteSlice(received)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Crypto
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Crypto"))
	return Go{classdb.Crypto(object)}
}

/*
Generates a [PackedByteArray] of cryptographically secure random bytes with given [param size].
*/
//go:nosplit
func (self class) GenerateRandomBytes(size gd.Int) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_generate_random_bytes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Generates an RSA [CryptoKey] that can be used for creating self-signed certificates and passed to [method StreamPeerTLS.accept_stream].
*/
//go:nosplit
func (self class) GenerateRsa(size gd.Int) gdclass.CryptoKey {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_generate_rsa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.CryptoKey{classdb.CryptoKey(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Generates a self-signed [X509Certificate] from the given [CryptoKey] and [param issuer_name]. The certificate validity will be defined by [param not_before] and [param not_after] (first valid date and last valid date). The [param issuer_name] must contain at least "CN=" (common name, i.e. the domain name), "O=" (organization, i.e. your company name), "C=" (country, i.e. 2 lettered ISO-3166 code of the country the organization is based in).
A small example to generate an RSA key and an X509 self-signed certificate.
[codeblocks]
[gdscript]
var crypto = Crypto.new()
# Generate 4096 bits RSA key.
var key = crypto.generate_rsa(4096)
# Generate self-signed certificate using the given key.
var cert = crypto.generate_self_signed_certificate(key, "CN=example.com,O=A Game Company,C=IT")
[/gdscript]
[csharp]
var crypto = new Crypto();
// Generate 4096 bits RSA key.
CryptoKey key = crypto.GenerateRsa(4096);
// Generate self-signed certificate using the given key.
X509Certificate cert = crypto.GenerateSelfSignedCertificate(key, "CN=mydomain.com,O=My Game Company,C=IT");
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GenerateSelfSignedCertificate(key gdclass.CryptoKey, issuer_name gd.String, not_before gd.String, not_after gd.String) gdclass.X509Certificate {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(key[0])[0])
	callframe.Arg(frame, discreet.Get(issuer_name))
	callframe.Arg(frame, discreet.Get(not_before))
	callframe.Arg(frame, discreet.Get(not_after))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_generate_self_signed_certificate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.X509Certificate{classdb.X509Certificate(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Sign a given [param hash] of type [param hash_type] with the provided private [param key].
*/
//go:nosplit
func (self class) Sign(hash_type classdb.HashingContextHashType, hash gd.PackedByteArray, key gdclass.CryptoKey) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, hash_type)
	callframe.Arg(frame, discreet.Get(hash))
	callframe.Arg(frame, discreet.Get(key[0])[0])
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_sign, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Verify that a given [param signature] for [param hash] of type [param hash_type] against the provided public [param key].
*/
//go:nosplit
func (self class) Verify(hash_type classdb.HashingContextHashType, hash gd.PackedByteArray, signature gd.PackedByteArray, key gdclass.CryptoKey) bool {
	var frame = callframe.New()
	callframe.Arg(frame, hash_type)
	callframe.Arg(frame, discreet.Get(hash))
	callframe.Arg(frame, discreet.Get(signature))
	callframe.Arg(frame, discreet.Get(key[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_verify, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Encrypt the given [param plaintext] with the provided public [param key].
[b]Note:[/b] The maximum size of accepted plaintext is limited by the key size.
*/
//go:nosplit
func (self class) Encrypt(key gdclass.CryptoKey, plaintext gd.PackedByteArray) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(key[0])[0])
	callframe.Arg(frame, discreet.Get(plaintext))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_encrypt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Decrypt the given [param ciphertext] with the provided private [param key].
[b]Note:[/b] The maximum size of accepted ciphertext is limited by the key size.
*/
//go:nosplit
func (self class) Decrypt(key gdclass.CryptoKey, ciphertext gd.PackedByteArray) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(key[0])[0])
	callframe.Arg(frame, discreet.Get(ciphertext))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_decrypt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Generates an [url=https://en.wikipedia.org/wiki/HMAC]HMAC[/url] digest of [param msg] using [param key]. The [param hash_type] parameter is the hashing algorithm that is used for the inner and outer hashes.
Currently, only [constant HashingContext.HASH_SHA256] and [constant HashingContext.HASH_SHA1] are supported.
*/
//go:nosplit
func (self class) HmacDigest(hash_type classdb.HashingContextHashType, key gd.PackedByteArray, msg gd.PackedByteArray) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, hash_type)
	callframe.Arg(frame, discreet.Get(key))
	callframe.Arg(frame, discreet.Get(msg))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_hmac_digest, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Compares two [PackedByteArray]s for equality without leaking timing information in order to prevent timing attacks.
See [url=https://paragonie.com/blog/2015/11/preventing-timing-attacks-on-string-comparison-with-double-hmac-strategy]this blog post[/url] for more information.
*/
//go:nosplit
func (self class) ConstantTimeCompare(trusted gd.PackedByteArray, received gd.PackedByteArray) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(trusted))
	callframe.Arg(frame, discreet.Get(received))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Crypto.Bind_constant_time_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCrypto() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCrypto() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("Crypto", func(ptr gd.Object) any { return classdb.Crypto(ptr) })}
