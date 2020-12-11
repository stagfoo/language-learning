local imageFile
local x
local y

function love.load()
    imageFile = love.graphics.newImage("images/1p.png")
    x = imageFile:getWidth()
    y = imageFile:getHeight()
end
function love.draw()
    love.window.setMode(x, y, {resizable=true})
    love.graphics.draw(imageFile)
end


function love.conf(t)
	t.console = true
end