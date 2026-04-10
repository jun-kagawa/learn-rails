class HelloController < ApplicationController
  def hello
    sleep 1
    res = {
      "hello": "world"
    }
    render json: res
  end
end
