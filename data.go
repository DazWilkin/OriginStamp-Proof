package main

var proofs = []struct {
	Hash string
	XML  []byte
}{{
	Hash: "2c5d36be542f8f0e7345d77753a5d7ea61a443ba6a9a86bb060332ad56dba38e",
	XML: []byte(`
<node value="69af555efcc31073fcb82977a5beb6fd84f20fe01f663cb1585c84945193f950" type="key">
	<left value="e4d9506f6a877a6f15c4e240e53327cc7c38ef37e610ebfc62ce674308b3263e" type="mesh">
		<left value="ee7536d1855cddf24dba36bfc04cdc1db522cef46e49c4d5aa5bde25648db388" type="mesh">
			<left value="9d7c106e64d5e6c40daacb4bcf4397912e31a78fe0e81b9fb4d79da08c7b2808" type="mesh">
				<left value="b918032b5cc7185ee2de8d38907e43b9ef9ffb61b5424ebf8a70643721c0af2f" type="mesh">
					<left value="4b2b83117d905892053e5f4a4d28316a0107b9a360fce4633361e489b7b1359e" type="mesh">
						<left value="10555d6d9b28b8cc551cd2dd99fb270630b2cf2f56f8d70a7893a64b4b57ecb5" type="mesh">
							<left value="c54c5350dba05d216175318dbfa6dd732d1adbdc92beaf320443f523f822f01d" type="mesh">
								<left value="c7ebc60a38854a4671ea786f14ceff0e7e183f75d6fbda3557292bae3b13f23d" type="mesh"/>
								<right value="a97c1a50c069d3397e9702f28ab38a420b537c6e942b8546860f8fb7a49c70e4" type="mesh">
									<left value="d1776b7c1b51ad1441bf765e194377ac9060f2402d6e91cee5502d8ca3a99546" type="mesh"/>
									<right value="25d8c347fb71a839c1274ec2ebe666b8dd4a3255403d991be39bb10016aa51f1" type="mesh">
										<left value="c006e869d7ba65ac4d5105888fee6a6bf6a3b38f23c6300137f003a3988238ae" type="mesh">
											<left value="097c397d862795e96abb576eb6bc76365f0235cfa1c938352045d1e62b347b1d" type="mesh"/>
											<right value="9407b62021fe6335430fc54090b90e6588ac38421309e279460335b5f94d67fb" type="mesh">
												<left value="369cb74a155479eecf22f65337e03a4ad5b935772e2d3348854f3d9d71923960" type="mesh"/>
												<right value="e220f15411b706ad10ff9990dc7f760ddfec2ebc61a11204fc5b36e140c17e0b" type="mesh">
													<left value="0d899a0e8f72691e3859377c4780a5b54c40e743d1486aeb1e76c038dab9e8ce" type="mesh"/>
													<right value="e40a25b8bd05b48675a02c8ad7e0c23b6d62385b0acaffbab63f95e39369ebc7" type="mesh">
														<left value="037e945cf8da5945acbcf2390c71a497c6edefdc364ada1f33d76a2b5f8b472b" type="hash"/>
														<right value="038df16cfbbeed2f029f00587bfcaa9920d4f7bcaeee60ee788d4a66179709c3" type="mesh"/>
													</right>
												</right>
											</right>
										</left>
										<right value="c3d44098be84f892ee1fe1c2c6caa52257eb96caaaa3f313edbafdc4a3ba6eae" type="mesh"/>
									</right>
								</right>
							</left>
							<right value="5b2e60ccb6907fae79d249bf7943b050c28b5fef84b9b44399f58461dd120c4a" type="mesh"/>
						</left>
						<right value="c47e581dd90c10aad2cb4121c1a94a3199f9788f4f2e5d0f200f8f167ae2517c" type="mesh"/>
					</left>
					<right value="dbdfc7af18a876d64d6fa2f0110e5dc6ead648f3e33f9e0163ceb267d47cd90d" type="mesh"/>
				</left>
				<right value="98f7dafd1f8c6ec84a555819d9cd05baed657e3973f86f3fb178dbd13c676bba" type="mesh"/>
			</left>
			<right value="36447318383768a406815764e128d404ca1482ec57cd995e12ac97962b33b73e" type="mesh"/>
		</left>
		<right value="6e7ca4d2ea8919396dfd66a7d57338cb0ce8bca76043b5a4745e463c2c9d9af8" type="mesh"/>
	</left>
	<right value="0f1301c53b6885f23037a81da41f450968a2b133f28b30d65d1fb431da74a9c1" type="mesh"/>
</node>
`),
}, {
	Hash: "037e945cf8da5945acbcf2390c71a497c6edefdc364ada1f33d76a2b5f8b472b",
	XML: []byte(`
<node value="5e92ec09501a5d39e251a151f84b5e2228312c445eb23b4e1de6360e27bad54b" type="key">
	<left value="7811e3130908fd2678eb2dd3928d245db7f3ed578c21f2ae4fbe680424dc735e" type="mesh">
		<left value="61f48c118002c0e7681e425a3b0e2396475fe0d037ebb0360231b95c2fd60c2f" type="mesh">
			<left value="54280d75669fb9f3ff976c66d03daefa02b691f1aa25558a85f86a5c60961c69" type="mesh">
				<left value="43e89c93c4247511a6eabfe24dfcd331453e04bdde13e688865bf73c3243280a" type="mesh"/>
				<right value="6c127084785466791190c1b4673635274d21e065b9ad9c70d246ef41ba960220" type="mesh">
					<left value="34789458c3010230cef61b6053626ee25f1d79762a76de3517ec5ac2e76ae2e9" type="mesh"/>
					<right value="f375cd66a5552b189a2bfe5f8be433236ff6829dea7bca710e611e98bc7baf0e" type="mesh">
						<left value="5e8d9839b88965971e8f8eb06d2d6bf2f72f75b16100ea175350179c6a82807e" type="mesh">
							<left value="55ed3f4915ab9f93e385f655a2c58c62a5fa2d0e7a4b741263102184e2b7a94f" type="mesh">
								<left value="c04c5bb97b65d009582feb32d88498880216aee8e3bda0c5a0f4ba12244aedb0" type="mesh"/>
								<right value="0bee8d27b7fd9129df599166751af1a9e82bf4e7ac91771f9c15ebf0cb05d74d" type="mesh">
									<left value="88369cb8594af97fb2ee465d5dc3a01e1e23924b8929faca6fdd7293f97e92be" type="mesh">
										<left value="d7ed34b6b74056df3cfb2836bfd55c4cd27b5237ad4a6641cf7003d21be76c88" type="mesh"/>
										<right value="af68501bb80cca4c0bc7d4df6aa2a0c347712e1f2cb1b687a0467e2e9aaf545e" type="mesh">
											<left value="debd8f8a4aa7af602faa9e8e08f5773c6decc29c9f65e80d70ee90a993170f69" type="mesh">
												<left value="8a477ae8ff3b601e709321b17f00117159a7efd861abc74286f7ee6479fe46ce" type="mesh">
													<left value="6a2368efa74c641f2f32d6be8f90a7b6ab5a3235d816faf4034d53d49ce5537c" type="mesh"/>
													<right value="2796b0a5986fe8f497057da457cdcf51ab9fc0ffb996445434bd08e2257b29e8" type="mesh">
														<left value="2c5838c92697fe3d5b8aa14e8f67f6fd2ff4c9e12d536332f57434ebcf0f0196" type="mesh"/>
														<right value="2c5d36be542f8f0e7345d77753a5d7ea61a443ba6a9a86bb060332ad56dba38e" type="hash"/>
													</right>
												</left>
												<right value="3d2466ee2746f818708f8f396649e0cc4d13738248e8ee051a767474ac2f87d3" type="mesh"/>
											</left>
											<right value="130921ae3e04ea13d237a9ed09e33c18574068e8db52e3d7d3d4855d7df3b3ce" type="mesh"/>
										</right>
									</left>
									<right value="8a5d691e559e95448c9aab2d92ba6d58ec9f814ee345e10aee81cf47790d7c2e" type="mesh"/>
								</right>
							</left>
							<right value="55562de568950252020816b6204cbc76820c7c9874c2a4672c99437d99068cf5" type="mesh"/>
						</left>
						<right value="f5fef74b0d216ca6221b024cc5f53d5aee7c9bc33f7289c529f465a0ff2969ea" type="mesh"/>
					</right>
				</right>
			</left>
			<right value="584d06e48bbfe9c5319495b7b3be9eaf7e11aa546ae777575e7e422aafa68fcd" type="mesh"/>
		</left>
		<right value="bf11a94cc0028fe2e0fe476d897da3417cc72777e523d6f0c030c5141bbeb75b" type="mesh"/>
	</left>
	<right value="a62eedb07080ce5a21ad26230bcd50ef37cac8cca43a2f1d946db2d1d47e1f94" type="mesh"/>
</node>	
`),
},
}
