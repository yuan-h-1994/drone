# frozen_string_literal: true

require_relative 'spec_helper'
require_relative '../hello_world'

describe 'Hello' do
  it '挨拶する' do
    expect(Hello.new.world).to eq('Hello World')
  end
end
