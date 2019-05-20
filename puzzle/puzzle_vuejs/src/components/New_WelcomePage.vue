<style scoped lang="less">
	.content {
		margin: 3em;
	}

	.cards {
		display: flex;
		.card {
			width: 352px;
      height: 538px;
      margin-left: 100px;
			background-color: white;
			border: 6px solid #131745;
      border-radius: 24px;
		}
	}

	.text {
		font-family: Fira Sans;
		font-style: normal;
		font-weight: 500;
		font-size: 24px;
		line-height: 29px;
		margin-top: 300px !important;

		color: #343d90;
  }
  
  .join-now {
    text-align: center;
    margin-top: 10px;
  }
</style>

<template >
	<div class="new-welcome-page">
		<div class="new-logo"></div>
		<div class="content">
			<div class="cards">
				<div class="card card-1">
					<div class="layer2">
						<div class="layer1">
							<div class="icon"></div>
							<div class="text">
								The contentos harmony puzzle game utilizes fast finality of the
								game to make it frictionless and the immutability
							</div>
						</div>
					</div>
				</div>

				<div class="card card-2">
					<div class="layer2">
						<div class="layer1">
							<div class="icon"></div>
							<div class="text">
								The contentos harmony puzzle game utilizes fast finality of the
								game to make it frictionless and the immutability
							</div>
						</div>
					</div>
				</div>

				<div class="card card-3">
					<div class="layer2">
						<div class="layer1">
							<div class="icon"></div>
							<div class="text">
								The contentos harmony puzzle game utilizes fast finality of the
								game to make it frictionless and the immutability
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
    <div class="join-now">
      <button class="btn-primary join-now" @click="$emit('join')">Join Now</button>
    </div>
	</div>
</template>

<script>
	export default {
		name: "NewWelcomePage",

		props: {},
		data() {
			return {};
		},
		mounted: function() {},
		methods: {}
	};
</script>
