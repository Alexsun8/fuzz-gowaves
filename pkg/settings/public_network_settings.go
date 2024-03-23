package settings

const mainnetSettingsJSON = `
{
  "features_voting_period": 5000,
  "votes_for_feature_activation": 4000,
  "preactivated_features": null,
  "double_features_periods_after_height": 810000,
  "sponsorship_single_activation_period": false,
  "generation_balance_depth_from_50_to_1000_after_height": 232000,
  "block_version_3_after_height": 795000,
  "reset_effective_balance_at_height": 462000,
  "stolen_aliases_window_time_start": 1522463241035,
  "stolen_aliases_window_time_end": 1530161445559,
  "reissue_bug_window_time_start": 1522463241035,
  "reissue_bug_window_time_end": 1530161445559,
  "allow_multiple_lease_cancel_until_time": 1492768800000,
  "allow_leased_balance_transfer_until_time": 1513357014002,
  "check_temp_negative_after_time": 1479168000000,
  "tx_changes_sorted_check_after_time": 1479416400000,
  "tx_from_future_check_after_time": 1479168000000,
  "unissued_asset_until_time": 1479416400000,
  "invalid_reissue_in_same_block_until_time": 1492768800000,
  "minimal_generating_balance_check_after_time": 1479168000000,
  "internal_invoke_payments_validation_after_height": 2959400,
  "internal_invoke_correct_fail_reject_behaviour_after_height": 2792473,
  "invoke_no_zero_payments_after_height": 0,
  "max_tx_time_back_offset": 7200000,
  "max_tx_time_forward_offset": 5400000,
  "address_scheme_character": 87,
  "average_block_delay_seconds": 60,
  "min_block_time": 15000,
  "delay_delta": 8,
  "max_base_target": 18446744073709551615,
  "block_reward_term": 100000,
  "block_reward_term_after_20": 50000,
  "initial_block_reward": 600000000,
  "block_reward_increment": 50000000,
  "block_reward_voting_period": 10000,
  "reward_addresses": [
    "3PEgG7eZHLFhcfsTSaYxgRhZsh4AxMvA4Ms",
    "3PFjHWuH6WXNJbwnfLHqNFBpwBS5dkYjTfv"
  ],
  "reward_addresses_after_21": [
    "3PEgG7eZHLFhcfsTSaYxgRhZsh4AxMvA4Ms"
  ],
  "min_xtn_buy_back_period": 100000,
  "min_update_asset_info_interval": 100000,
  "light_node_block_fields_absence_interval": 1000,
  "type": 0,
  "genesis": {
    "version": 1,
    "timestamp": 1460678400000,
    "reference": "67rpwLCuS5DGA8KGZXKsVQ7dnPb9goRLoKfgGbLfQg9WoLUgNY77E2jT11fem3coV9nAkguBACzrU1iyZM4B8roQ",
    "desiredReward": 0,
    "nxt-consensus": {
      "base-target": 153722867,
      "generation-signature": "11111111111111111111111111111111"
    },
    "transactionBlockLength": 283,
    "transactionCount": 6,
    "signature": "FSH8eAAzZNqnG8xgTZtz5xuLqXySsXgAjmFEC25hXMbEufiGjqWPnGCZFt6gLiVLJny16ipxRNAkkzjjhqTjBE2",
    "height": 0,
    "transactions": [
      {
        "type": 1,
        "id": "2DVtfgXjpMeFf2PQCqvwxAiaGbiDsxDjSdNQkc5JQ74eWxjWFYgwvqzC4dn7iB1AhuM32WxEiVi1SGijsBtYQwn8",
        "signature": "2DVtfgXjpMeFf2PQCqvwxAiaGbiDsxDjSdNQkc5JQ74eWxjWFYgwvqzC4dn7iB1AhuM32WxEiVi1SGijsBtYQwn8",
        "timestamp": 1465742577614,
        "recipient": "3PAWwWa6GbwcJaFzwqXQN5KQm7H96Y7SHTQ",
        "amount": 9999999500000000
      },
      {
        "type": 1,
        "id": "2TsxPS216SsZJAiep7HrjZ3stHERVkeZWjMPFcvMotrdGpFa6UCCmoFiBGNizx83Ks8DnP3qdwtJ8WFcN9J4exa3",
        "signature": "2TsxPS216SsZJAiep7HrjZ3stHERVkeZWjMPFcvMotrdGpFa6UCCmoFiBGNizx83Ks8DnP3qdwtJ8WFcN9J4exa3",
        "timestamp": 1465742577614,
        "recipient": "3P8JdJGYc7vaLu4UXUZc1iRLdzrkGtdCyJM",
        "amount": 100000000
      },
      {
        "type": 1,
        "id": "3gF8LFjhnZdgEVjP7P6o1rvwapqdgxn7GCykCo8boEQRwxCufhrgqXwdYKEg29jyPWthLF5cFyYcKbAeFvhtRNTc",
        "signature": "3gF8LFjhnZdgEVjP7P6o1rvwapqdgxn7GCykCo8boEQRwxCufhrgqXwdYKEg29jyPWthLF5cFyYcKbAeFvhtRNTc",
        "timestamp": 1465742577614,
        "recipient": "3PAGPDPqnGkyhcihyjMHe9v36Y4hkAh9yDy",
        "amount": 100000000
      },
      {
        "type": 1,
        "id": "5hjSPLDyqic7otvtTJgVv73H3o6GxgTBqFMTY2PqAFzw2GHAnoQddC4EgWWFrAiYrtPadMBUkoepnwFHV1yR6u6g",
        "signature": "5hjSPLDyqic7otvtTJgVv73H3o6GxgTBqFMTY2PqAFzw2GHAnoQddC4EgWWFrAiYrtPadMBUkoepnwFHV1yR6u6g",
        "timestamp": 1465742577614,
        "recipient": "3P9o3ZYwtHkaU1KxsKkFjJqJKS3dLHLC9oF",
        "amount": 100000000
      },
      {
        "type": 1,
        "id": "ivP1MzTd28yuhJPkJsiurn2rH2hovXqxr7ybHZWoRGUYKazkfaL9MYoTUym4sFgwW7WB5V252QfeFTsM6Uiz3DM",
        "signature": "ivP1MzTd28yuhJPkJsiurn2rH2hovXqxr7ybHZWoRGUYKazkfaL9MYoTUym4sFgwW7WB5V252QfeFTsM6Uiz3DM",
        "timestamp": 1465742577614,
        "recipient": "3PJaDyprvekvPXPuAtxrapacuDJopgJRaU3",
        "amount": 100000000
      },
      {
        "type": 1,
        "id": "29gnRjk8urzqc9kvqaxAfr6niQTuTZnq7LXDAbd77nydHkvrTA4oepoMLsiPkJ8wj2SeFB5KXASSPmbScvBbfLiV",
        "signature": "29gnRjk8urzqc9kvqaxAfr6niQTuTZnq7LXDAbd77nydHkvrTA4oepoMLsiPkJ8wj2SeFB5KXASSPmbScvBbfLiV",
        "timestamp": 1465742577614,
        "recipient": "3PBWXDFUc86N2EQxKJmW8eFco65xTyMZx6J",
        "amount": 100000000
      }
    ]
  }
}`

const testnetSettingsJSON = `
{
  "features_voting_period": 3000,
  "votes_for_feature_activation": 2700,
  "preactivated_features": null,
  "double_features_periods_after_height": 18446744073709551615,
  "sponsorship_single_activation_period": false,
  "generation_balance_depth_from_50_to_1000_after_height": 0,
  "block_version_3_after_height": 161700,
  "reset_effective_balance_at_height": 51500,
  "stolen_aliases_window_time_start": 0,
  "stolen_aliases_window_time_end": 0,
  "reissue_bug_window_time_start": 1520411086003,
  "reissue_bug_window_time_end": 1523096218005,
  "allow_multiple_lease_cancel_until_time": 1492560000000,
  "allow_leased_balance_transfer_until_time": 1508230496004,
  "check_temp_negative_after_time": 1477958400000,
  "tx_changes_sorted_check_after_time": 1479416400000,
  "tx_from_future_check_after_time": 1478100000000,
  "unissued_asset_until_time": 1479416400000,
  "invalid_reissue_in_same_block_until_time": 1492560000000,
  "minimal_generating_balance_check_after_time": 0,
  "internal_invoke_payments_validation_after_height": 1698800,
  "internal_invoke_correct_fail_reject_behaviour_after_height": 1727461,
  "invoke_no_zero_payments_after_height": 0,
  "max_tx_time_back_offset": 7200000,
  "max_tx_time_forward_offset": 5400000,
  "address_scheme_character": 84,
  "average_block_delay_seconds": 60,
  "min_block_time": 15000,
  "delay_delta": 8,
  "max_base_target": 18446744073709551615,
  "block_reward_term": 100000,
  "block_reward_term_after_20": 50000,
  "initial_block_reward": 600000000,
  "block_reward_increment": 50000000,
  "block_reward_voting_period": 10000,
  "reward_addresses": [
    "3Myb6G8DkdBb8YcZzhrky65HrmiNuac3kvS",
    "3N13KQpdY3UU7JkWUBD9kN7t7xuUgeyYMTT"
  ],
  "reward_addresses_after_21": [
    "3Myb6G8DkdBb8YcZzhrky65HrmiNuac3kvS"
  ],
  "min_xtn_buy_back_period": 2000,
  "min_update_asset_info_interval": 100000,
  "light_node_block_fields_absence_interval": 1000,
  "type": 1,
  "genesis": {
    "version": 1,
    "timestamp": 1460678400000,
    "reference": "67rpwLCuS5DGA8KGZXKsVQ7dnPb9goRLoKfgGbLfQg9WoLUgNY77E2jT11fem3coV9nAkguBACzrU1iyZM4B8roQ",
    "desiredReward": 0,
    "nxt-consensus": {
      "base-target": 153722867,
      "generation-signature": "11111111111111111111111111111111"
    },
    "transactionBlockLength": 236,
    "transactionCount": 5,
    "signature": "5uqnLK3Z9eiot6FyYBfwUnbyid3abicQbAZjz38GQ1Q8XigQMxTK4C1zNkqS1SVw7FqSidbZKxWAKLVoEsp4nNqa",
    "height": 0,
    "transactions": [
      {
        "type": 1,
        "id": "5G66c9GPn2egiM4bQBBF3gCkHS8sQZupRvWCpWKWGQTRRbtqdtZJ5Mt29exbHTDZW2RWygVKZ3oBNg4RwezN7wmA",
        "signature": "5G66c9GPn2egiM4bQBBF3gCkHS8sQZupRvWCpWKWGQTRRbtqdtZJ5Mt29exbHTDZW2RWygVKZ3oBNg4RwezN7wmA",
        "timestamp": 1478000000000,
        "recipient": "3My3KZgFQ3CrVHgz6vGRt8687sH4oAA1qp8",
        "amount": 400000000000000
      },
      {
        "type": 1,
        "id": "3zpi4i5SeCoaiCBn1iuTUvCc5aahvtabqXBTrCXy1Y3ujUbJo56VVv6n4HQtcwiFapvg3BKV6stb5QkxsBrudTKZ",
        "signature": "3zpi4i5SeCoaiCBn1iuTUvCc5aahvtabqXBTrCXy1Y3ujUbJo56VVv6n4HQtcwiFapvg3BKV6stb5QkxsBrudTKZ",
        "timestamp": 1478000000000,
        "recipient": "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
        "amount": 200000000000000
      },
      {
        "type": 1,
        "id": "3obfFPvsWXv2RyMYxjTT7owYGcpSGuSAm8fQVXeX5wErWYsgNSPPnQoFVV6nzuwm3RwGCbm8dfgvqwK9S8fVMpye",
        "signature": "3obfFPvsWXv2RyMYxjTT7owYGcpSGuSAm8fQVXeX5wErWYsgNSPPnQoFVV6nzuwm3RwGCbm8dfgvqwK9S8fVMpye",
        "timestamp": 1478000000000,
        "recipient": "3N5GRqzDBhjVXnCn44baHcz2GoZy5qLxtTh",
        "amount": 200000000000000
      },
      {
        "type": 1,
        "id": "3TdE9G7V7fwED35981aGsWFM6aesxSS4W1XPfEx6p5xacwHLu7Kvf67Wzg73kgyU9gSFp1KsmPWqkFhaaR2S1fhp",
        "signature": "3TdE9G7V7fwED35981aGsWFM6aesxSS4W1XPfEx6p5xacwHLu7Kvf67Wzg73kgyU9gSFp1KsmPWqkFhaaR2S1fhp",
        "timestamp": 1478000000000,
        "recipient": "3NCBMxgdghg4tUhEEffSXy11L6hUi6fcBpd",
        "amount": 200000000000000
      },
      {
        "type": 1,
        "id": "4hTrr7fqkujsGSH8AFN1qw7fJdfmKgwzoq3ByCCJwduHkgZPQZe1KgzG6oPBZXMuNr5ZQ6ErDSTiz2KGtxtkHpA5",
        "signature": "4hTrr7fqkujsGSH8AFN1qw7fJdfmKgwzoq3ByCCJwduHkgZPQZe1KgzG6oPBZXMuNr5ZQ6ErDSTiz2KGtxtkHpA5",
        "timestamp": 1478000000000,
        "recipient": "3N18z4B8kyyQ96PhN5eyhCAbg4j49CgwZJx",
        "amount": 9000000000000000
      }
    ]
  }
}`

const stagenetSettingsJSON = `
{
  "features_voting_period": 100,
  "votes_for_feature_activation": 40,
  "preactivated_features": [
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
    11,
    12,
    13
  ],
  "double_features_periods_after_height": 1000000000,
  "sponsorship_single_activation_period": false,
  "generation_balance_depth_from_50_to_1000_after_height": 0,
  "block_version_3_after_height": 0,
  "reset_effective_balance_at_height": 0,
  "stolen_aliases_window_time_start": 0,
  "stolen_aliases_window_time_end": 0,
  "reissue_bug_window_time_start": 0,
  "reissue_bug_window_time_end": 0,
  "allow_multiple_lease_cancel_until_time": 0,
  "allow_leased_balance_transfer_until_time": 0,
  "check_temp_negative_after_time": 0,
  "tx_changes_sorted_check_after_time": 0,
  "tx_from_future_check_after_time": 0,
  "unissued_asset_until_time": 0,
  "invalid_reissue_in_same_block_until_time": 0,
  "minimal_generating_balance_check_after_time": 0,
  "internal_invoke_payments_validation_after_height": 966180,
  "internal_invoke_correct_fail_reject_behaviour_after_height": 390000,
  "invoke_no_zero_payments_after_height": 1317000,
  "max_tx_time_back_offset": 7200000,
  "max_tx_time_forward_offset": 5400000,
  "address_scheme_character": 83,
  "average_block_delay_seconds": 60,
  "min_block_time": 15000,
  "delay_delta": 8,
  "max_base_target": 18446744073709551615,
  "block_reward_term": 100000,
  "block_reward_term_after_20": 50000,
  "initial_block_reward": 600000000,
  "block_reward_increment": 50000000,
  "block_reward_voting_period": 10000,
  "reward_addresses": [
    "3MaFVH1vTv18FjBRugSRebx259D7xtRh9ic",
    "3MbhiRiLFLJ1EVKNP9npRszcLLQDjwnFfZM"
  ],
  "reward_addresses_after_21": [
    "3MaFVH1vTv18FjBRugSRebx259D7xtRh9ic"
  ],
  "min_xtn_buy_back_period": 1000,
  "min_update_asset_info_interval": 10,
  "light_node_block_fields_absence_interval": 1000,
  "type": 2,
  "genesis": {
    "version": 1,
    "timestamp": 1561705836768,
    "reference": "67rpwLCuS5DGA8KGZXKsVQ7dnPb9goRLoKfgGbLfQg9WoLUgNY77E2jT11fem3coV9nAkguBACzrU1iyZM4B8roQ",
    "desiredReward": 0,
    "nxt-consensus": {
      "base-target": 5000,
      "generation-signature": "11111111111111111111111111111111"
    },
    "transactionBlockLength": 48,
    "transactionCount": 1,
    "signature": "2EaaguFPgrJ1bbMAFrPw2bi6i7kqjgvxsFj8YGqrKR7hT54ZvwmzZ3LHMm4qR7i7QB5cacp8XdkLMJyvjFkt8VgN",
    "height": 0,
    "transactions": [
      {
        "type": 1,
        "id": "4jE5MLtQhHcdtj2U4emsN9FiBs7cLvdx7BwHSX2rVShW46Q5S8e61mXpigfm26Y4igNpAc4MhfeuLzv6pa3CySHL",
        "signature": "4jE5MLtQhHcdtj2U4emsN9FiBs7cLvdx7BwHSX2rVShW46Q5S8e61mXpigfm26Y4igNpAc4MhfeuLzv6pa3CySHL",
        "timestamp": 1561705836768,
        "recipient": "3Mi63XiwniEj6mTC557pxdRDddtpj7fZMMw",
        "amount": 10000000000000000
      }
    ]
  }
}`
