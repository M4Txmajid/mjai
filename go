-- ╔═══════════════════════════════════════════════════════════════╗
-- ║   MUSLIM DEVS - PREMIUM EDITION + BRAINROT ESP                ║
-- ║   Multi-Map Support + Full ESP System + SAFE TP               ║
-- ║   MOBILE + PC SUPPORT                                          ║
-- ╚═══════════════════════════════════════════════════════════════╝

local Players = game:GetService("Players")
local TweenService = game:GetService("TweenService")
local RunService = game:GetService("RunService")
local UserInputService = game:GetService("UserInputService")
local CoreGui = game:GetService("CoreGui")

local LP = Players.LocalPlayer
local Camera = workspace.CurrentCamera

-- ════════════════════════════════════════
-- // CLEANUP OLD INSTANCES
-- ════════════════════════════════════════
pcall(function()
    local oldGui = LP:WaitForChild("PlayerGui"):FindFirstChild("MuslimDevs")
    if oldGui then oldGui:Destroy() end
end)
pcall(function()
    if CoreGui:FindFirstChild("BrainrotESPSelector") then
        CoreGui:FindFirstChild("BrainrotESPSelector"):Destroy()
    end
end)
pcall(function()
    if getgenv().BrainrotESP_RenderConnection then
        getgenv().BrainrotESP_RenderConnection:Disconnect()
    end
end)
pcall(function()
    if getgenv().BrainrotESP_Objects then
        for _, data in pairs(getgenv().BrainrotESP_Objects) do
            if data.esp then
                for _, line in pairs(data.esp.BoxLines or {}) do pcall(function() line:Remove() end) end
                pcall(function() data.esp.Name:Remove() end)
                pcall(function() data.esp.Rarity:Remove() end)
                pcall(function() data.esp.Distance:Remove() end)
                pcall(function() data.esp.Tracer:Remove() end)
            end
        end
    end
end)
pcall(function()
    for _, part in pairs(workspace:GetChildren()) do
        if part.Name:find("CoinScanner") or part.Name:find("DoomScanner") then part:Destroy() end
    end
end)

-- ════════════════════════════════════════
-- // MOBILE DETECTION
-- ════════════════════════════════════════
local IsMobile = UserInputService.TouchEnabled and not UserInputService.KeyboardEnabled
local ScreenSize = Camera.ViewportSize

-- Responsive sizing
local Scale = IsMobile and math.clamp(ScreenSize.X / 500, 0.7, 1) or 1
local MainWidth = IsMobile and math.clamp(math.floor(ScreenSize.X * 0.92), 300, 400) or 440
local MainHeight = IsMobile and math.clamp(math.floor(ScreenSize.Y * 0.72), 380, 500) or 540
local TabWidth = IsMobile and 50 or 58
local HeaderHeight = IsMobile and 52 or 60
local BtnHeight = IsMobile and 40 or 34
local ToggleH = IsMobile and 44 or 38
local CardPad = IsMobile and 10 or 12
local FontMult = IsMobile and 1.1 or 1
local GapRowH = IsMobile and 36 or 30
local TabIconSize = IsMobile and 18 or 16
local TabLabelSize = IsMobile and 9 or 8
local TabBtnH = IsMobile and 48 or 42

-- ════════════════════════════════════════
-- // VARIABLES (ORIGINAL SCRIPT)
-- ════════════════════════════════════════
local AllGaps = {}
local LastTP = nil
local LastTPIndex = nil
local AutoReTP = false
local AutoCollect = false
local AutoCollectSpeed = 0.3
local MoneyCollectCount = 0
local AutoRebirth = false
local RebirthSpeed = 1
local RebirthCount = 0
local AutoUpgrade1 = false
local AutoUpgrade10 = false
local UpgradeSpeed = 0.5
local Upgrade1Count = 0
local Upgrade10Count = 0
local AutoUpgradeBase = false
local UpgradeBaseSpeed = 0.5
local UpgradeBaseCount = 0
local MyBase = nil

-- Coin Scanner
local AutoCoinScan = false
local ScanRadius = 50
local ScanSpeed = 2
local ScannerParts = {}
local RadioactiveCollected = 0
local DoomCollected = 0
local TotalCoinsCollected = 0
local scanAngle = 0
local pulseTime = 0
local scannerConnection = nil

-- Wall & Grab
local WallsCleared = false
local InstantGrabEnabled = false
local GrabConnection = nil
local BrainrotsFixed = 0
local InstantLuckyBlockEnabled = false
local LuckyBlockConnection = nil
local LuckyBlocksFixed = 0

-- Map Detection
local CurrentMap = "Unknown"

-- TP Variables
local SuccessDistance = 5
local TeleportAttempts = 0
local TeleportSuccesses = 0
local IsAtTarget = false
local MainTPLoop = nil

-- Safe TP Variables
local SpawnCFrame = CFrame.new(72.9999924, -2.99998093, -6.34033713e-05)
local UndergroundY = -42
local SafeTPSpeed = 200
local SafeTPVertSpeed = 150
local IsSafeTPing = false
local SafeTPPhase = ""
local NoclipEnabled = false
local NoclipConnection = nil
local SafeTPCancel = false



-- ════════════════════════════════════════
-- // ANTI-CHEAT SCANNER VARIABLES
-- ════════════════════════════════════════
local AC_ScanResults = {
    Remotes = {},
    Scripts = {},
    Modules = {},
    Misc = {},
    Deleted = {},
}
local AC_TotalFound = 0
local AC_TotalDisabled = 0
local AC_TotalHooked = 0
local AC_ScanComplete = false
local AC_IsScanning = false
local AC_HookedRemotes = {}
local AC_OriginalNamecallHook = nil

local AC_Keywords = {
    "anticheat", "anti_cheat", "anti-cheat", "anticheats",
    "antiexploit", "anti_exploit", "anti-exploit",
    "antiteleport", "anti_teleport", "anti_tp", "antitp",
    "antispeed", "anti_speed", "antinoclip", "anti_noclip",
    "antifly", "anti_fly",
    "security", "securitycheck", "securitysystem",
    "guard", "guardian", "watchdog", "sentinel",
    "protection", "protector", "shield",
    "integrity", "integritycheck",
    "detect", "detector", "detection", "detected",
    "monitor", "monitoring", "watchman",
    "checker", "checking", "checkplayer",
    "scanner", "scanning",
    "kick", "kickplayer", "kickuser",
    "ban", "banplayer", "banuser", "bancheck",
    "punish", "punishment", "penalty",
    "flag", "flagged", "flagplayer",
    "report", "reportplayer", "autoreport",
    "verify", "verification", "validate", "validation",
    "authenticate", "auth", "legit",
    "sanity", "sanitycheck",
    "poscheck", "positioncheck", "checkpos", "checkposition",
    "speedcheck", "checkspeed", "velocitycheck",
    "teleportcheck", "tpcheck", "checkteleport", "checktp",
    "movementcheck", "checkmovement",
    "distancecheck", "checkdistance",
    "noclipcheck", "checknoclip",
    "flycheck", "checkfly", "heightcheck", "checkheight",
    "exploit", "exploiter", "exploitcheck",
    "cheat", "cheater", "cheating", "cheatcheck",
    "hack", "hacker", "hacking", "hackcheck",
    "abuse", "abuser",
    "inject", "injector", "injection",
    "heartbeat_ac", "acheartbeat",
    "clientcheck", "servercheck",
    "ping_check", "pingcheck",
    "clientverify", "clientvalidate",
    "ac_", "_ac",
}

local AC_Patterns = {
    "anti%w*cheat", "anti%w*exploit", "anti%w*tp",
    "anti%w*teleport", "anti%w*speed", "anti%w*fly", "anti%w*noclip",
    "check%w*pos", "check%w*speed", "check%w*teleport",
    "check%w*movement", "check%w*player",
    "verify%w*client", "validate%w*pos",
    "detect%w*exploit", "detect%w*cheat", "detect%w*hack",
    "kick%w*player", "ban%w*player", "flag%w*player",
    "security%w*check", "integrity%w*check", "sanity%w*check",
}
-- ════════════════════════════════════════
-- // ESP SETTINGS & VARIABLES
-- ════════════════════════════════════════
local ESPConfig = {
    Enabled = false,
    MaxDistance = 2000,
    ShowBox = true,
    ShowName = true,
    ShowDistance = true,
    ShowTracer = true,
    ShowRarity = true,
    TextSize = 14,
    BoxThickness = 1.5,
    TracerThickness = 1.2,
}

local RarityColors = {
    ["Common"]    = Color3.fromRGB(180, 180, 180),
    ["Uncommon"]  = Color3.fromRGB(80, 200, 80),
    ["Rare"]      = Color3.fromRGB(70, 130, 255),
    ["Epic"]      = Color3.fromRGB(180, 60, 255),
    ["Legendary"] = Color3.fromRGB(255, 170, 0),
    ["Mythic"]    = Color3.fromRGB(255, 50, 50),
    ["Divine"]    = Color3.fromRGB(255, 215, 0),
    ["Secret"]    = Color3.fromRGB(255, 100, 200),
    ["Default"]   = Color3.fromRGB(255, 85, 255),
}

local SelectedFolders = {}
local SelectedBrainrots = {}
local ESPObjects = {}
getgenv().BrainrotESP_Objects = ESPObjects

-- ════════════════════════════════════════
-- // THEME
-- ════════════════════════════════════════
local Theme = {
    Primary = Color3.fromRGB(30, 144, 255),
    PrimaryDark = Color3.fromRGB(0, 100, 200),
    PrimaryLight = Color3.fromRGB(100, 180, 255),
    Accent = Color3.fromRGB(255, 215, 0),
    Bg = Color3.fromRGB(10, 10, 20),
    BgLight = Color3.fromRGB(18, 18, 35),
    BgLighter = Color3.fromRGB(28, 28, 50),
    BgHover = Color3.fromRGB(38, 38, 65),
    Text = Color3.fromRGB(255, 255, 255),
    TextDim = Color3.fromRGB(180, 180, 200),
    TextMuted = Color3.fromRGB(120, 120, 140),
    Green = Color3.fromRGB(80, 220, 120),
    Red = Color3.fromRGB(255, 80, 80),
    Purple = Color3.fromRGB(150, 100, 255),
    Blue = Color3.fromRGB(80, 150, 255),
    Discord = Color3.fromRGB(88, 101, 242),
    Orange = Color3.fromRGB(255, 140, 50),
    Yellow = Color3.fromRGB(255, 220, 50),
    Cyan = Color3.fromRGB(50, 220, 220),
    Radioactive = Color3.fromRGB(80, 255, 80),
    Pink = Color3.fromRGB(255, 100, 200),
    Gold = Color3.fromRGB(255, 200, 50),
    VIP = Color3.fromRGB(255, 215, 0),
    Doom = Color3.fromRGB(255, 50, 50),
    ESP = Color3.fromRGB(255, 85, 255),
}

-- ════════════════════════════════════════
-- // UI HELPERS
-- ════════════════════════════════════════
local function Tween(obj, props, dur, style, dir)
    local t = TweenService:Create(obj, TweenInfo.new(dur or 0.3, style or Enum.EasingStyle.Quint, dir or Enum.EasingDirection.Out), props)
    t:Play()
    return t
end

local function Corner(parent, radius)
    local c = Instance.new("UICorner")
    c.CornerRadius = UDim.new(0, radius or 10)
    c.Parent = parent
    return c
end

local function Stroke(parent, color, thick, trans)
    local s = Instance.new("UIStroke")
    s.Color = color or Theme.PrimaryLight
    s.Thickness = thick or 1.5
    s.Transparency = trans or 0.3
    s.Parent = parent
    return s
end

local function Ripple(button)
    button.ClipsDescendants = true
    button.MouseButton1Click:Connect(function()
        local r = Instance.new("Frame")
        r.Size = UDim2.new(0, 0, 0, 0)
        r.Position = UDim2.new(0.5, 0, 0.5, 0)
        r.AnchorPoint = Vector2.new(0.5, 0.5)
        r.BackgroundColor3 = Color3.new(1, 1, 1)
        r.BackgroundTransparency = 0.7
        r.BorderSizePixel = 0
        r.ZIndex = button.ZIndex + 1
        r.Parent = button
        Corner(r, 100)
        local s = math.max(button.AbsoluteSize.X, button.AbsoluteSize.Y) * 2.5
        Tween(r, {Size = UDim2.new(0, s, 0, s), BackgroundTransparency = 1}, 0.5)
        task.delay(0.5, function() if r then r:Destroy() end end)
    end)
end

local function Hover(btn, normal, hover)
    if IsMobile then return end
    btn.MouseEnter:Connect(function() Tween(btn, {BackgroundColor3 = hover}, 0.12) end)
    btn.MouseLeave:Connect(function() Tween(btn, {BackgroundColor3 = normal}, 0.12) end)
end

local function FS(size)
    return math.floor(size * FontMult)
end

local function CreateRainbowText(parent, text, size, position, textSize)
    local container = Instance.new("Frame")
    container.Name = "RainbowTextContainer"
    container.Size = size
    container.Position = position
    container.BackgroundTransparency = 1
    container.ZIndex = 11
    container.Parent = parent

    local textLabels = {}
    local textLength = #text

    for i = 1, textLength do
        local char = text:sub(i, i)
        local label = Instance.new("TextLabel")
        label.Name = "Char" .. i
        label.Size = UDim2.new(1/textLength, 0, 1, 0)
        label.Position = UDim2.new((i-1)/textLength, 0, 0, 0)
        label.BackgroundTransparency = 1
        label.Text = char
        label.TextColor3 = Color3.fromHSV((i-1)/textLength, 1, 1)
        label.TextSize = FS(textSize or 18)
        label.Font = Enum.Font.GothamBlack
        label.TextXAlignment = Enum.TextXAlignment.Center
        label.ZIndex = 11
        label.Parent = container
        table.insert(textLabels, label)
    end

    task.spawn(function()
        local offset = 0
        while container and container.Parent do
            offset = offset + 0.02
            for i, label in ipairs(textLabels) do
                if label and label.Parent then
                    local hue = ((i-1)/textLength + offset) % 1
                    label.TextColor3 = Color3.fromHSV(hue, 1, 1)
                end
            end
            task.wait(0.03)
        end
    end)

    return container, textLabels
end

-- ════════════════════════════════════════
-- // ESP UTILITY FUNCTIONS
-- ════════════════════════════════════════
local function GetRarityColor(rarity)
    return RarityColors[rarity] or RarityColors["Default"]
end

local function GetRealBrainrotName(brainrot)
    local renderBrainrot = brainrot:FindFirstChild("RenderBrainrot")
    if renderBrainrot then
        return renderBrainrot.Name
    end

    for _, child in pairs(brainrot:GetChildren()) do
        if child:IsA("Model") and child.Name ~= "HumanoidRootPart" and child.Name ~= "Head" then
            return child.Name
        end
    end

    local nameValue = brainrot:FindFirstChild("BrainrotName") or brainrot:FindFirstChild("DisplayName")
    if nameValue and nameValue:IsA("StringValue") and nameValue.Value ~= "" then
        return nameValue.Value
    end

    local attrName = brainrot:GetAttribute("BrainrotName") or brainrot:GetAttribute("DisplayName")
    if attrName then return attrName end

    return brainrot.Name
end

local function IsESPSelected(folder, brainrotName)
    if SelectedFolders[folder] == true then return true end
    if SelectedBrainrots[folder .. "/" .. brainrotName] == true then return true end
    return false
end

local function NothingESPSelected()
    for _, v in pairs(SelectedFolders) do
        if v then return false end
    end
    for _, v in pairs(SelectedBrainrots) do
        if v then return false end
    end
    return true
end

local function ClearAllESPSelections()
    for key, _ in pairs(SelectedFolders) do
        SelectedFolders[key] = nil
    end
    for key, _ in pairs(SelectedBrainrots) do
        SelectedBrainrots[key] = nil
    end
    for brainrot, data in pairs(ESPObjects) do
        if data.esp then
            for _, line in pairs(data.esp.BoxLines or {}) do 
                pcall(function() line.Visible = false end) 
            end
            pcall(function() data.esp.Name.Visible = false end)
            pcall(function() data.esp.Rarity.Visible = false end)
            pcall(function() data.esp.Distance.Visible = false end)
            pcall(function() data.esp.Tracer.Visible = false end)
        end
    end
end

local function ScanBrainrotFolders()
    local ActiveBrainrots = workspace:FindFirstChild("ActiveBrainrots")
    if not ActiveBrainrots then return {} end

    local structure = {}

    for _, child in pairs(ActiveBrainrots:GetChildren()) do
        if child:IsA("Folder") then
            structure[child.Name] = {}
            for _, brainrot in pairs(child:GetChildren()) do
                if brainrot:IsA("Model") then
                    table.insert(structure[child.Name], {
                        modelName = brainrot.Name,
                        realName = GetRealBrainrotName(brainrot)
                    })
                end
            end
        elseif child:IsA("Model") then
            if not structure["Uncategorized"] then
                structure["Uncategorized"] = {}
            end
            table.insert(structure["Uncategorized"], {
                modelName = child.Name,
                realName = GetRealBrainrotName(child)
            })
        end
    end

    return structure
end

-- ════════════════════════════════════════
-- // ESP DRAWING SYSTEM
-- ════════════════════════════════════════
local function CreateESP()
    local esp = {}

    esp.BoxLines = {}
    for i = 1, 4 do
        local line = Drawing.new("Line")
        line.Visible = false
        line.Thickness = ESPConfig.BoxThickness
        table.insert(esp.BoxLines, line)
    end

    esp.Name = Drawing.new("Text")
    esp.Name.Visible = false
    esp.Name.Size = ESPConfig.TextSize
    esp.Name.Center = true
    esp.Name.Outline = true
    esp.Name.OutlineColor = Color3.fromRGB(0, 0, 0)
    esp.Name.Font = Drawing.Fonts.Plex

    esp.Rarity = Drawing.new("Text")
    esp.Rarity.Visible = false
    esp.Rarity.Size = ESPConfig.TextSize - 2
    esp.Rarity.Center = true
    esp.Rarity.Outline = true
    esp.Rarity.OutlineColor = Color3.fromRGB(0, 0, 0)
    esp.Rarity.Font = Drawing.Fonts.Plex

    esp.Distance = Drawing.new("Text")
    esp.Distance.Visible = false
    esp.Distance.Size = ESPConfig.TextSize - 2
    esp.Distance.Center = true
    esp.Distance.Outline = true
    esp.Distance.OutlineColor = Color3.fromRGB(0, 0, 0)
    esp.Distance.Font = Drawing.Fonts.Plex
    esp.Distance.Color = Color3.fromRGB(200, 200, 200)

    esp.Tracer = Drawing.new("Line")
    esp.Tracer.Visible = false
    esp.Tracer.Thickness = ESPConfig.TracerThickness

    return esp
end

local function RemoveESP(esp)
    if not esp then return end
    for _, line in pairs(esp.BoxLines or {}) do pcall(function() line:Remove() end) end
    pcall(function() esp.Name:Remove() end)
    pcall(function() esp.Rarity:Remove() end)
    pcall(function() esp.Distance:Remove() end)
    pcall(function() esp.Tracer:Remove() end)
end

local function HideESP(esp)
    if not esp then return end
    for _, line in pairs(esp.BoxLines) do line.Visible = false end
    esp.Name.Visible = false
    esp.Rarity.Visible = false
    esp.Distance.Visible = false
    esp.Tracer.Visible = false
end

local function GetBrainrotInfo(brainrot)
    local part = brainrot:FindFirstChild("HumanoidRootPart")
        or brainrot:FindFirstChild("Head")
        or brainrot:FindFirstChildWhichIsA("BasePart")

    if not part then
        local render = brainrot:FindFirstChild("RenderBrainrot")
        if render then
            part = render:FindFirstChild("HumanoidRootPart")
                or render:FindFirstChild("Head")
                or render:FindFirstChildWhichIsA("BasePart")
        end
    end

    if not part then return nil end

    local success, cf, size = pcall(function()
        return brainrot:GetBoundingBox()
    end)

    if not success then
        cf = part.CFrame
        size = part.Size
    end

    return {
        Position = cf.Position,
        Size = size,
    }
end

local function UpdateESP(brainrot, esp, folderName, realName)
    if not ESPConfig.Enabled then HideESP(esp) return end
    if not brainrot or not brainrot.Parent then HideESP(esp) return end

    local nothingIsSelected = NothingESPSelected()
    if nothingIsSelected then
        HideESP(esp)
        return
    end
    
    if not IsESPSelected(folderName, realName) then
        HideESP(esp)
        return
    end

    local data = GetBrainrotInfo(brainrot)
    if not data then HideESP(esp) return end

    local char = LP.Character
    if not char or not char:FindFirstChild("HumanoidRootPart") then HideESP(esp) return end

    local dist = (char.HumanoidRootPart.Position - data.Position).Magnitude
    if dist > ESPConfig.MaxDistance then HideESP(esp) return end

    local screenPos, onScreen = Camera:WorldToViewportPoint(data.Position)
    local color = GetRarityColor(folderName)

    if not onScreen then
        for _, line in pairs(esp.BoxLines) do line.Visible = false end
        esp.Name.Visible = false
        esp.Rarity.Visible = false
        esp.Distance.Visible = false

        if ESPConfig.ShowTracer then
            local bottom = Vector2.new(Camera.ViewportSize.X / 2, Camera.ViewportSize.Y)
            local cx = math.clamp(screenPos.X, 0, Camera.ViewportSize.X)
            local cy = math.clamp(screenPos.Y, 0, Camera.ViewportSize.Y)
            esp.Tracer.From = bottom
            esp.Tracer.To = Vector2.new(cx, cy)
            esp.Tracer.Color = color
            esp.Tracer.Visible = true
        else
            esp.Tracer.Visible = false
        end
        return
    end

    local topW = Camera:WorldToViewportPoint(data.Position + Vector3.new(0, data.Size.Y / 2, 0))
    local botW = Camera:WorldToViewportPoint(data.Position - Vector3.new(0, data.Size.Y / 2, 0))
    local boxH = math.abs(botW.Y - topW.Y)
    local boxW = boxH * 0.6
    local bx = screenPos.X - boxW / 2
    local by = topW.Y

    local tl = Vector2.new(bx, by)
    local tr = Vector2.new(bx + boxW, by)
    local bl = Vector2.new(bx, by + boxH)
    local br = Vector2.new(bx + boxW, by + boxH)

    if ESPConfig.ShowBox then
        esp.BoxLines[1].From = tl  esp.BoxLines[1].To = tr
        esp.BoxLines[2].From = tr  esp.BoxLines[2].To = br
        esp.BoxLines[3].From = br  esp.BoxLines[3].To = bl
        esp.BoxLines[4].From = bl  esp.BoxLines[4].To = tl
        for _, line in pairs(esp.BoxLines) do
            line.Color = color
            line.Thickness = ESPConfig.BoxThickness
            line.Visible = true
        end
    else
        for _, line in pairs(esp.BoxLines) do line.Visible = false end
    end

    local yOffset = by - ESPConfig.TextSize - 4
    if ESPConfig.ShowName then
        esp.Name.Text = realName
        esp.Name.Position = Vector2.new(screenPos.X, yOffset)
        esp.Name.Color = Color3.fromRGB(255, 255, 255)
        esp.Name.Visible = true
        yOffset = yOffset - ESPConfig.TextSize
    else
        esp.Name.Visible = false
    end

    if ESPConfig.ShowRarity then
        esp.Rarity.Text = "[" .. folderName .. "]"
        esp.Rarity.Position = Vector2.new(screenPos.X, yOffset)
        esp.Rarity.Color = color
        esp.Rarity.Visible = true
    else
        esp.Rarity.Visible = false
    end

    if ESPConfig.ShowDistance then
        esp.Distance.Text = string.format("[%d studs]", math.floor(dist))
        esp.Distance.Position = Vector2.new(screenPos.X, by + boxH + 2)
        esp.Distance.Visible = true
    else
        esp.Distance.Visible = false
    end

    if ESPConfig.ShowTracer then
        local bottom = Vector2.new(Camera.ViewportSize.X / 2, Camera.ViewportSize.Y)
        esp.Tracer.From = bottom
        esp.Tracer.To = Vector2.new(screenPos.X, by + boxH)
        esp.Tracer.Color = color
        esp.Tracer.Visible = true
    else
        esp.Tracer.Visible = false
    end
end

-- ════════════════════════════════════════
-- // ESP BRAINROT TRACKING
-- ════════════════════════════════════════
local function AddBrainrotESP(brainrot, folderName)
    if ESPObjects[brainrot] then return end
    local realName = GetRealBrainrotName(brainrot)
    ESPObjects[brainrot] = {
        esp = CreateESP(),
        folder = folderName,
        realName = realName
    }
end

local function RemoveBrainrotESP(brainrot)
    if not ESPObjects[brainrot] then return end
    RemoveESP(ESPObjects[brainrot].esp)
    ESPObjects[brainrot] = nil
end

local function TrackESPFolder(folder)
    local folderName = folder.Name
    for _, brainrot in pairs(folder:GetChildren()) do
        if brainrot:IsA("Model") then
            AddBrainrotESP(brainrot, folderName)
        end
    end
    folder.ChildAdded:Connect(function(brainrot)
        if brainrot:IsA("Model") then
            task.wait(0.2)
            AddBrainrotESP(brainrot, folderName)
        end
    end)
    folder.ChildRemoved:Connect(function(brainrot)
        RemoveBrainrotESP(brainrot)
    end)
end

local function InitESPTracking()
    local ActiveBrainrots = workspace:FindFirstChild("ActiveBrainrots")
    if not ActiveBrainrots then return end

    for _, child in pairs(ActiveBrainrots:GetChildren()) do
        if child:IsA("Folder") then
            TrackESPFolder(child)
        elseif child:IsA("Model") then
            AddBrainrotESP(child, "Uncategorized")
        end
    end

    ActiveBrainrots.ChildAdded:Connect(function(child)
        task.wait(0.2)
        if child:IsA("Folder") then
            TrackESPFolder(child)
        elseif child:IsA("Model") then
            AddBrainrotESP(child, "Uncategorized")
        end
    end)

    ActiveBrainrots.ChildRemoved:Connect(function(child)
        if child:IsA("Folder") then
            for _, b in pairs(child:GetChildren()) do RemoveBrainrotESP(b) end
        else
            RemoveBrainrotESP(child)
        end
    end)
end

-- ════════════════════════════════════════
-- // GAME LOGIC - MULTI-MAP SUPPORT
-- ════════════════════════════════════════
local function RefreshGaps()
    AllGaps = {}

    local mapNames = {
        "DefaultMap_SharedInstances",
        "DoomMap_SharedInstances",
        "ValentinesMap_SharedInstances",
    }

    local directMapNames = {
        "ValentinesMap",
    }

    for _, mapName in ipairs(mapNames) do
        local mapFolder = workspace:FindFirstChild(mapName)
        if mapFolder then
            local gapsFolder = mapFolder:FindFirstChild("Gaps")
            if gapsFolder then
                for _, gap in ipairs(gapsFolder:GetChildren()) do
                    local mud = gap:FindFirstChild("Mud")
                    if not mud then
                        for _, child in ipairs(gap:GetDescendants()) do
                            if child.Name:lower():find("mud") then mud = child break end
                        end
                    end
                    if mud then
                        local mudPart = mud:IsA("BasePart") and mud or mud:FindFirstChildWhichIsA("BasePart", true)
                        if mudPart then
                            table.insert(AllGaps, {
                                Name = gap.Name,
                                Mud = mudPart,
                                Map = mapName:gsub("_SharedInstances", "")
                            })
                        end
                    end
                end
            end
        end
    end

    for _, mapName in ipairs(directMapNames) do
        local alreadyFound = false
        for _, gap in ipairs(AllGaps) do
            if gap.Map == mapName then alreadyFound = true break end
        end
        if not alreadyFound then
            local mapFolder = workspace:FindFirstChild(mapName)
            if mapFolder then
                local gapsFolder = mapFolder:FindFirstChild("Gaps")
                if gapsFolder then
                    for _, gap in ipairs(gapsFolder:GetChildren()) do
                        local mud = gap:FindFirstChild("Mud")
                        if not mud then
                            for _, child in ipairs(gap:GetDescendants()) do
                                if child.Name:lower():find("mud") then mud = child break end
                            end
                        end
                        if mud then
                            local mudPart = mud:IsA("BasePart") and mud or mud:FindFirstChildWhichIsA("BasePart", true)
                            if mudPart then
                                table.insert(AllGaps, {
                                    Name = gap.Name,
                                    Mud = mudPart,
                                    Map = mapName
                                })
                            end
                        end
                    end
                end
            end
        end
    end

    table.sort(AllGaps, function(a, b)
        if a.Map ~= b.Map then return a.Map < b.Map end
        return (tonumber(a.Name:match("%d+")) or 0) < (tonumber(b.Name:match("%d+")) or 0)
    end)

    if workspace:FindFirstChild("DoomMap_SharedInstances") then
        CurrentMap = "DoomMap"
    elseif workspace:FindFirstChild("ValentinesMap_SharedInstances") or workspace:FindFirstChild("ValentinesMap") then
        CurrentMap = "ValentinesMap"
    elseif workspace:FindFirstChild("DefaultMap_SharedInstances") then
        CurrentMap = "DefaultMap"
    else
        CurrentMap = "Unknown"
    end

    return #AllGaps
end

-- ════════════════════════════════════════
-- // TP SYSTEM
-- ════════════════════════════════════════
local function IsAlive()
    local char = LP.Character
    if not char then return false end
    local humanoid = char:FindFirstChild("Humanoid")
    if not humanoid then return false end
    if humanoid.Health <= 0 then return false end
    local hrp = char:FindFirstChild("HumanoidRootPart")
    if not hrp then return false end
    return true
end

local function GetHRP()
    local char = LP.Character
    if not char then return nil end
    return char:FindFirstChild("HumanoidRootPart")
end

local function GetDistance(targetCFrame)
    local hrp = GetHRP()
    if not hrp then return math.huge end
    if not targetCFrame then return math.huge end
    return (hrp.Position - targetCFrame.Position).Magnitude
end

local function DoInstantTP(targetCFrame)
    local hrp = GetHRP()
    if not hrp then return false end
    if not targetCFrame then return false end
    pcall(function()
        hrp.Velocity = Vector3.zero
        hrp.AssemblyLinearVelocity = Vector3.zero
        hrp.AssemblyAngularVelocity = Vector3.zero
    end)
    hrp.CFrame = targetCFrame
    pcall(function()
        hrp.Velocity = Vector3.zero
        hrp.AssemblyLinearVelocity = Vector3.zero
        hrp.AssemblyAngularVelocity = Vector3.zero
    end)
    TeleportAttempts = TeleportAttempts + 1
    return true
end

local function StartMainTPLoop()
    if MainTPLoop then return end
    MainTPLoop = RunService.RenderStepped:Connect(function()
        if not AutoReTP then IsAtTarget = false return end
        if not LastTP then IsAtTarget = false return end
        if not IsAlive() then IsAtTarget = false return end
        local distance = GetDistance(LastTP)
        if distance <= SuccessDistance then
            if not IsAtTarget then
                TeleportSuccesses = TeleportSuccesses + 1
                IsAtTarget = true
            end
        else
            IsAtTarget = false
            DoInstantTP(LastTP)
        end
    end)
end

local function StopMainTPLoop()
    if MainTPLoop then MainTPLoop:Disconnect() MainTPLoop = nil end
    IsAtTarget = false
end

local function TPToMud(index)
    if index < 1 or index > #AllGaps then return end
    local mud = AllGaps[index].Mud
    local target = CFrame.new(mud.Position + Vector3.new(0, mud.Size.Y / 2 + 3, 0))
    LastTP = target
    LastTPIndex = index
    IsAtTarget = false
    DoInstantTP(target)
end

-- ════════════════════════════════════════
-- // NOCLIP SYSTEM
-- ════════════════════════════════════════
local function EnableNoclip()
    if NoclipConnection then return end
    NoclipEnabled = true
    NoclipConnection = RunService.Stepped:Connect(function()
        if not NoclipEnabled then return end
        local char = LP.Character
        if not char then return end
        for _, part in pairs(char:GetDescendants()) do
            if part:IsA("BasePart") then
                part.CanCollide = false
            end
        end
    end)
end

local function DisableNoclip()
    NoclipEnabled = false
    if NoclipConnection then
        NoclipConnection:Disconnect()
        NoclipConnection = nil
    end
    local char = LP.Character
    if char then
        for _, part in pairs(char:GetDescendants()) do
            if part:IsA("BasePart") and (part.Name == "HumanoidRootPart" or part.Name == "Head" or part.Name:find("Torso") or part.Name:find("Leg") or part.Name:find("Arm")) then
                part.CanCollide = true
            end
        end
    end
end

-- ════════════════════════════════════════
-- // SAFE TP TO SPAWN
-- ════════════════════════════════════════
local function SafeTPToSpawn(statusLabel, phaseLabel, progressBar)
    if IsSafeTPing then return end
    if not IsAlive() then
        if statusLabel then statusLabel.Text = "❌ You're dead!" statusLabel.TextColor3 = Theme.Red end
        return
    end

    IsSafeTPing = true
    SafeTPCancel = false

    local hrp = GetHRP()
    if not hrp then
        IsSafeTPing = false
        return
    end

    local startPos = hrp.Position
    local targetPos = SpawnCFrame.Position
    local distDown = math.abs(startPos.Y - UndergroundY)
    local distHorizontal = (Vector3.new(startPos.X, 0, startPos.Z) - Vector3.new(targetPos.X, 0, targetPos.Z)).Magnitude
    local distUp = math.abs(UndergroundY - targetPos.Y)
    local totalDist = distDown + distHorizontal + distUp

    EnableNoclip()

    if statusLabel then
        statusLabel.Text = "🛡️ SAFE TP ACTIVE"
        statusLabel.TextColor3 = Theme.Cyan
    end

    pcall(function()
        hrp.Velocity = Vector3.zero
        hrp.AssemblyLinearVelocity = Vector3.zero
        hrp.AssemblyAngularVelocity = Vector3.zero
    end)

    -- PHASE 1: DIVE
    SafeTPPhase = "⬇️ Diving Underground"
    if phaseLabel then
        phaseLabel.Text = "⬇️ Phase 1/3: Diving to Y=" .. UndergroundY
        phaseLabel.TextColor3 = Theme.Yellow
    end

    local undergroundTarget = CFrame.new(startPos.X, UndergroundY, startPos.Z)

    while IsSafeTPing and not SafeTPCancel do
        hrp = GetHRP()
        if not hrp or not IsAlive() then break end

        local currentPos = hrp.Position
        local target = undergroundTarget.Position
        local direction = target - currentPos
        local distance = direction.Magnitude

        if distance < 1.5 then
            hrp.CFrame = undergroundTarget
            break
        end

        local step = direction.Unit * math.min(SafeTPVertSpeed * (1/60), distance)
        hrp.CFrame = CFrame.new(currentPos + step)

        pcall(function()
            hrp.Velocity = Vector3.zero
            hrp.AssemblyLinearVelocity = Vector3.zero
            hrp.AssemblyAngularVelocity = Vector3.zero
        end)

        if progressBar then
            progressBar.Size = UDim2.new(math.clamp((distDown - distance) / totalDist, 0, 1), 0, 1, 0)
        end

        RunService.RenderStepped:Wait()
    end

    if SafeTPCancel then
        DisableNoclip()
        IsSafeTPing = false
        SafeTPPhase = ""
        if statusLabel then statusLabel.Text = "⛔ CANCELLED" statusLabel.TextColor3 = Theme.Red end
        if phaseLabel then phaseLabel.Text = "" end
        return
    end

    task.wait(0.1)

    -- PHASE 2: HORIZONTAL
    SafeTPPhase = "➡️ Moving Underground"
    if phaseLabel then
        phaseLabel.Text = string.format("➡️ Phase 2/3: Underground → Spawn (%.0f studs)", distHorizontal)
        phaseLabel.TextColor3 = Theme.Cyan
    end

    local horizontalTarget = CFrame.new(targetPos.X, UndergroundY, targetPos.Z)

    while IsSafeTPing and not SafeTPCancel do
        hrp = GetHRP()
        if not hrp or not IsAlive() then break end

        local currentPos = hrp.Position
        local target = horizontalTarget.Position
        local direction = target - currentPos
        local distance = direction.Magnitude

        if distance < 1.5 then
            hrp.CFrame = horizontalTarget
            break
        end

        local step = direction.Unit * math.min(SafeTPSpeed * (1/60), distance)
        hrp.CFrame = CFrame.new(currentPos + step)

        pcall(function()
            hrp.Velocity = Vector3.zero
            hrp.AssemblyLinearVelocity = Vector3.zero
            hrp.AssemblyAngularVelocity = Vector3.zero
        end)

        local currentTraveled = distDown + (distHorizontal - distance)
        if progressBar then
            progressBar.Size = UDim2.new(math.clamp(currentTraveled / totalDist, 0, 1), 0, 1, 0)
        end

        if phaseLabel then
            phaseLabel.Text = string.format("➡️ Phase 2/3: %.0f studs remaining", distance)
        end

        RunService.RenderStepped:Wait()
    end

    if SafeTPCancel then
        DisableNoclip()
        IsSafeTPing = false
        SafeTPPhase = ""
        if statusLabel then statusLabel.Text = "⛔ CANCELLED" statusLabel.TextColor3 = Theme.Red end
        if phaseLabel then phaseLabel.Text = "" end
        return
    end

    task.wait(0.1)

    -- PHASE 3: RISE
    SafeTPPhase = "⬆️ Rising to Spawn"
    if phaseLabel then
        phaseLabel.Text = "⬆️ Phase 3/3: Rising to spawn surface"
        phaseLabel.TextColor3 = Theme.Green
    end

    while IsSafeTPing and not SafeTPCancel do
        hrp = GetHRP()
        if not hrp or not IsAlive() then break end

        local currentPos = hrp.Position
        local target = SpawnCFrame.Position
        local direction = target - currentPos
        local distance = direction.Magnitude

        if distance < 1.5 then
            hrp.CFrame = SpawnCFrame
            break
        end

        local step = direction.Unit * math.min(SafeTPVertSpeed * (1/60), distance)
        hrp.CFrame = CFrame.new(currentPos + step)

        pcall(function()
            hrp.Velocity = Vector3.zero
            hrp.AssemblyLinearVelocity = Vector3.zero
            hrp.AssemblyAngularVelocity = Vector3.zero
        end)

        if progressBar then
            progressBar.Size = UDim2.new(math.clamp((totalDist - distance) / totalDist, 0, 0.99), 0, 1, 0)
        end

        RunService.RenderStepped:Wait()
    end

    task.wait(0.3)
    DisableNoclip()

    hrp = GetHRP()
    if hrp then
        hrp.CFrame = SpawnCFrame
        pcall(function()
            hrp.Velocity = Vector3.zero
            hrp.AssemblyLinearVelocity = Vector3.zero
        end)
    end

    IsSafeTPing = false
    SafeTPPhase = ""

    if progressBar then
        progressBar.Size = UDim2.new(1, 0, 1, 0)
        progressBar.BackgroundColor3 = Theme.Green
    end

    if statusLabel then
        statusLabel.Text = "✅ ARRIVED SAFELY AT SPAWN!"
        statusLabel.TextColor3 = Theme.Green
    end
    if phaseLabel then
        phaseLabel.Text = "🏁 Safe & Sound! No waves touched you."
        phaseLabel.TextColor3 = Theme.Green
    end
end

-- ════════════════════════════════════════
-- // BASE & AUTO FUNCTIONS
-- ════════════════════════════════════════
local function FindMyBase()
    if MyBase and MyBase.Parent then return MyBase end
    local basesFolder = workspace:FindFirstChild("Bases")
    if not basesFolder then return nil end
    local char = LP.Character
    if not char or not char:FindFirstChild("HumanoidRootPart") then return nil end
    local hrp = char.HumanoidRootPart
    local best, bestDist = nil, math.huge
    for _, base in ipairs(basesFolder:GetChildren()) do
        for _, part in ipairs(base:GetDescendants()) do
            if part:IsA("BasePart") then
                local d = (hrp.Position - part.Position).Magnitude
                if d < bestDist then bestDist = d best = base end
            end
        end
    end
    MyBase = best
    return MyBase
end

local function CollectSlots()
    if not MyBase or not MyBase.Parent then return end
    local char = LP.Character
    if not char then return end
    local hrp = char:FindFirstChild("HumanoidRootPart")
    if not hrp then return end
    local search = MyBase:FindFirstChild("Slots") or MyBase
    for _, desc in ipairs(search:GetDescendants()) do
        if desc:IsA("BasePart") and desc:FindFirstChild("TouchInterest") then
            pcall(function()
                if firetouchinterest then
                    firetouchinterest(hrp, desc, 0)
                    firetouchinterest(hrp, desc, 1)
                end
            end)
        end
    end
    MoneyCollectCount = MoneyCollectCount + 1
end

task.spawn(function()
    while true do
        if AutoCollect and MyBase then pcall(CollectSlots) end
        task.wait(AutoCollectSpeed or 0.3)
    end
end)

task.spawn(function()
    while true do
        if AutoRebirth then
            pcall(function()
                local result = game:GetService("ReplicatedStorage").RemoteFunctions.Rebirth:InvokeServer()
                if result then RebirthCount = RebirthCount + 1 end
            end)
        end
        task.wait(RebirthSpeed)
    end
end)

task.spawn(function()
    while true do
        if AutoUpgrade1 then
            pcall(function()
                game:GetService("ReplicatedStorage").RemoteFunctions.UpgradeSpeed:InvokeServer(1)
                Upgrade1Count = Upgrade1Count + 1
            end)
        end
        task.wait(UpgradeSpeed)
    end
end)

task.spawn(function()
    while true do
        if AutoUpgrade10 then
            pcall(function()
                game:GetService("ReplicatedStorage").RemoteFunctions.UpgradeSpeed:InvokeServer(10)
                Upgrade10Count = Upgrade10Count + 1
            end)
        end
        task.wait(UpgradeSpeed)
    end
end)

task.spawn(function()
    while true do
        if AutoUpgradeBase then
            pcall(function()
                game:GetService("ReplicatedStorage").Shared.Remotes.Networking["RE/Plots/PlotUpgradeBase"]:FireServer()
                UpgradeBaseCount = UpgradeBaseCount + 1
            end)
        end
        task.wait(UpgradeBaseSpeed)
    end
end)

-- ════════════════════════════════════════
-- // COIN SCANNER
-- ════════════════════════════════════════
local function CreateScannerVisuals()
    for _, part in pairs(ScannerParts) do
        if part and part.Parent then part:Destroy() end
    end
    ScannerParts = {}
    local scanner = Instance.new("Part")
    scanner.Name = "CoinScanner_Main"
    scanner.Shape = Enum.PartType.Cylinder
    scanner.Size = Vector3.new(0.5, ScanRadius * 2, ScanRadius * 2)
    scanner.Anchored = true
    scanner.CanCollide = false
    scanner.Transparency = 0.85
    scanner.Color = Theme.Cyan
    scanner.Material = Enum.Material.Neon
    scanner.CastShadow = false
    scanner.Parent = workspace
    table.insert(ScannerParts, scanner)

    local scanLine = Instance.new("Part")
    scanLine.Name = "CoinScanner_Line"
    scanLine.Size = Vector3.new(ScanRadius, 0.4, 0.4)
    scanLine.Anchored = true
    scanLine.CanCollide = false
    scanLine.Transparency = 0.3
    scanLine.Color = Theme.Yellow
    scanLine.Material = Enum.Material.Neon
    scanLine.CastShadow = false
    scanLine.Parent = workspace
    table.insert(ScannerParts, scanLine)

    local scanLine2 = Instance.new("Part")
    scanLine2.Name = "CoinScanner_Line2"
    scanLine2.Size = Vector3.new(ScanRadius, 0.3, 0.3)
    scanLine2.Anchored = true
    scanLine2.CanCollide = false
    scanLine2.Transparency = 0.5
    scanLine2.Color = Theme.Radioactive
    scanLine2.Material = Enum.Material.Neon
    scanLine2.CastShadow = false
    scanLine2.Parent = workspace
    table.insert(ScannerParts, scanLine2)

    return scanner, scanLine, scanLine2
end

local function DestroyScannerVisuals()
    for _, part in pairs(ScannerParts) do
        if part and part.Parent then part:Destroy() end
    end
    ScannerParts = {}
end

local function CollectCoinPart(part, hrp, coinType)
    pcall(function()
        if part:FindFirstChild("TouchInterest") then
            firetouchinterest(hrp, part, 0)
            task.wait()
            firetouchinterest(hrp, part, 1)
            TotalCoinsCollected = TotalCoinsCollected + 1
            if coinType == "radioactive" then RadioactiveCollected = RadioactiveCollected + 1
            else DoomCollected = DoomCollected + 1 end
            return
        end
        for _, child in pairs(part:GetDescendants()) do
            if child:IsA("TouchTransmitter") or child.Name == "TouchInterest" then
                firetouchinterest(hrp, child.Parent, 0)
                task.wait()
                firetouchinterest(hrp, child.Parent, 1)
                TotalCoinsCollected = TotalCoinsCollected + 1
                if coinType == "radioactive" then RadioactiveCollected = RadioactiveCollected + 1
                else DoomCollected = DoomCollected + 1 end
                return
            end
        end
    end)
end

local function ScanAndCollectAllCoins(hrp)
    local rootPos = hrp.Position
    local eventParts = workspace:FindFirstChild("EventParts")
    if eventParts then
        local radioFolder = eventParts:FindFirstChild("RadioactiveCoinsFolder")
        if radioFolder then
            for _, coin in pairs(radioFolder:GetChildren()) do
                local coinPart = coin:IsA("BasePart") and coin or (coin:IsA("Model") and (coin.PrimaryPart or coin:FindFirstChildWhichIsA("BasePart")))
                if coinPart then
                    local ok, dist = pcall(function() return (coinPart.Position - rootPos).Magnitude end)
                    if ok and dist <= ScanRadius then CollectCoinPart(coinPart, hrp, "radioactive") end
                end
            end
        end
    end
    local doomParts = workspace:FindFirstChild("DoomEventParts")
    if doomParts then
        for _, obj in pairs(doomParts:GetDescendants()) do
            if obj:IsA("BasePart") then
                local ok, dist = pcall(function() return (obj.Position - rootPos).Magnitude end)
                if ok and dist <= ScanRadius then
                    if obj:FindFirstChild("TouchInterest") or obj.Name:lower():find("coin") or obj.Name:lower():find("token") then
                        CollectCoinPart(obj, hrp, "doom")
                    end
                end
            end
        end
    end
end

local function StartCoinScanner()
    if scannerConnection then return end
    local scanner, scanLine, scanLine2 = CreateScannerVisuals()
    scannerConnection = RunService.RenderStepped:Connect(function(dt)
        if not AutoCoinScan then
            DestroyScannerVisuals()
            if scannerConnection then scannerConnection:Disconnect() scannerConnection = nil end
            return
        end
        local char = LP.Character
        if not char then return end
        local hrp = char:FindFirstChild("HumanoidRootPart")
        if not hrp then return end
        if not scanner or not scanner.Parent then scanner, scanLine, scanLine2 = CreateScannerVisuals() end
        local rootPos = hrp.Position
        local rootCFrame = CFrame.new(rootPos)
        scanAngle = scanAngle + (dt * ScanSpeed * math.pi * 2)
        pulseTime = pulseTime + dt
        local pulse = math.sin(pulseTime * 3) * 0.08 + 0.92
        if scanner and scanner.Parent then
            scanner.CFrame = rootCFrame * CFrame.Angles(0, 0, math.rad(90))
            scanner.Size = Vector3.new(0.5, ScanRadius * 2 * pulse, ScanRadius * 2 * pulse)
            scanner.Color = Color3.fromHSV((pulseTime * 0.1) % 1, 0.7, 1)
        end
        if scanLine and scanLine.Parent then
            scanLine.CFrame = rootCFrame * CFrame.Angles(0, scanAngle, 0) * CFrame.new(ScanRadius / 2, 0, 0)
        end
        if scanLine2 and scanLine2.Parent then
            scanLine2.CFrame = rootCFrame * CFrame.Angles(0, scanAngle + math.pi, 0) * CFrame.new(ScanRadius / 2, 0, 0)
        end
    end)
end

local function StopCoinScanner()
    if scannerConnection then scannerConnection:Disconnect() scannerConnection = nil end
    DestroyScannerVisuals()
end

task.spawn(function()
    while true do
        if AutoCoinScan then
            local char = LP.Character
            if char then
                local hrp = char:FindFirstChild("HumanoidRootPart")
                if hrp then pcall(function() ScanAndCollectAllCoins(hrp) end) end
            end
        end
        task.wait(0.15)
    end
end)

-- ════════════════════════════════════════
-- // INSTANT GRAB FUNCTIONS
-- ════════════════════════════════════════
local function fixAllBrainrots()
    local activeBrainrots = workspace:FindFirstChild("ActiveBrainrots")
    if not activeBrainrots then return 0 end
    local count = 0
    for _, rarity in pairs(activeBrainrots:GetChildren()) do
        for _, brainrot in pairs(rarity:GetChildren()) do
            local root = brainrot:FindFirstChild("Root")
            if root then
                local takePrompt = root:FindFirstChild("TakePrompt")
                if takePrompt and takePrompt:IsA("ProximityPrompt") and takePrompt.HoldDuration > 0 then
                    takePrompt.HoldDuration = 0
                    count = count + 1
                end
            end
        end
    end
    BrainrotsFixed = BrainrotsFixed + count
    return count
end

local function startInstantGrab()
    fixAllBrainrots()
    local activeBrainrots = workspace:FindFirstChild("ActiveBrainrots")
    if activeBrainrots then
        GrabConnection = activeBrainrots.DescendantAdded:Connect(function(desc)
            if InstantGrabEnabled and desc:IsA("ProximityPrompt") and desc.Name == "TakePrompt" then
                task.wait()
                desc.HoldDuration = 0
                BrainrotsFixed = BrainrotsFixed + 1
            end
        end)
    end
    task.spawn(function()
        while InstantGrabEnabled do pcall(fixAllBrainrots) task.wait(2) end
    end)
end

local function stopInstantGrab()
    if GrabConnection then GrabConnection:Disconnect() GrabConnection = nil end
end

local function fixAllLuckyBlocks()
    local activeLuckyBlocks = workspace:FindFirstChild("ActiveLuckyBlocks")
    if not activeLuckyBlocks then return 0 end
    local count = 0
    for _, luckyBlock in pairs(activeLuckyBlocks:GetDescendants()) do
        if luckyBlock:IsA("Model") then
            local rootPart = luckyBlock:FindFirstChild("RootPart")
            if rootPart then
                local pp = rootPart:FindFirstChild("ProximityPrompt")
                if pp and pp:IsA("ProximityPrompt") and pp.HoldDuration > 0 then
                    pp.HoldDuration = 0
                    count = count + 1
                end
            end
        end
    end
    LuckyBlocksFixed = LuckyBlocksFixed + count
    return count
end

local function startInstantLuckyBlock()
    fixAllLuckyBlocks()
    local activeLuckyBlocks = workspace:FindFirstChild("ActiveLuckyBlocks")
    if activeLuckyBlocks then
        LuckyBlockConnection = activeLuckyBlocks.DescendantAdded:Connect(function(desc)
            if InstantLuckyBlockEnabled and desc:IsA("ProximityPrompt") then
                task.wait()
                desc.HoldDuration = 0
                LuckyBlocksFixed = LuckyBlocksFixed + 1
            end
        end)
    end
    task.spawn(function()
        while InstantLuckyBlockEnabled do pcall(fixAllLuckyBlocks) task.wait(2) end
    end)
end

local function stopInstantLuckyBlock()
    if LuckyBlockConnection then LuckyBlockConnection:Disconnect() LuckyBlockConnection = nil end
end

-- ════════════════════════════════════════
-- // ANTI-CHEAT SCANNER FUNCTIONS
-- ════════════════════════════════════════
local function AC_IsACName(name)
    if not name or type(name) ~= "string" then return false, "" end
    local lowerName = name:lower()
    for _, keyword in ipairs(AC_Keywords) do
        if lowerName:find(keyword, 1, true) then return true, keyword end
    end
    for _, pattern in ipairs(AC_Patterns) do
        if lowerName:find(pattern) then return true, pattern end
    end
    return false, ""
end

local function AC_GetThreatLevel(name)
    local lowerName = (name or ""):lower()
    local highThreats = {"anticheat", "antiexploit", "kick", "ban", "punish", "antiteleport", "antitp", "anti_tp", "antispeed", "antinoclip", "antifly"}
    for _, t in ipairs(highThreats) do
        if lowerName:find(t, 1, true) then return "🔴 HIGH", Theme.Red end
    end
    local medThreats = {"detect", "monitor", "check", "verify", "validate", "security", "guard", "flag", "report", "exploit", "cheat", "hack"}
    for _, t in ipairs(medThreats) do
        if lowerName:find(t, 1, true) then return "🟡 MEDIUM", Theme.Yellow end
    end
    local lowThreats = {"position", "speed", "movement", "distance", "height", "velocity", "teleport", "heartbeat", "integrity", "sanity"}
    for _, t in ipairs(lowThreats) do
        if lowerName:find(t, 1, true) then return "🟠 LOW", Theme.Orange end
    end
    return "⚪ UNKNOWN", Theme.TextMuted
end

local function AC_GetCategory(item)
    if item:IsA("RemoteEvent") then return "🔴 RemoteEvent"
    elseif item:IsA("RemoteFunction") then return "🟠 RemoteFunction"
    elseif item:IsA("BindableEvent") then return "🟡 BindableEvent"
    elseif item:IsA("BindableFunction") then return "🟡 BindableFunction"
    elseif item:IsA("LocalScript") then return "📜 LocalScript"
    elseif item:IsA("ModuleScript") then return "📦 ModuleScript"
    elseif item:IsA("Script") then return "⚙️ Script"
    else return "❓ " .. item.ClassName end
end

local function AC_ScanLocation(location, locationName, results, statusCallback)
    if not location then return end
    pcall(function()
        for _, item in pairs(location:GetDescendants()) do
            local isAC, matchedKeyword = AC_IsACName(item.Name)
            if isAC then
                local threatLevel, threatColor = AC_GetThreatLevel(item.Name)
                local category = AC_GetCategory(item)
                local entry = {
                    Instance = item,
                    Name = item.Name,
                    ClassName = item.ClassName,
                    Path = item:GetFullName(),
                    Location = locationName,
                    MatchedKeyword = matchedKeyword,
                    ThreatLevel = threatLevel,
                    ThreatColor = threatColor,
                    Category = category,
                    Disabled = false,
                    Hooked = false,
                }
                if item:IsA("RemoteEvent") or item:IsA("RemoteFunction") or item:IsA("BindableEvent") or item:IsA("BindableFunction") then
                    table.insert(results.Remotes, entry)
                elseif item:IsA("LocalScript") then
                    table.insert(results.Scripts, entry)
                elseif item:IsA("ModuleScript") then
                    table.insert(results.Modules, entry)
                else
                    table.insert(results.Misc, entry)
                end
                AC_TotalFound = AC_TotalFound + 1
                if statusCallback then statusCallback("🔍 Found: " .. item.Name .. " in " .. locationName) end
            end
        end
    end)
end

local function AC_ScanForObfuscated(results, statusCallback)
    local function scanDesc(parent)
        pcall(function()
            for _, child in pairs(parent:GetDescendants()) do
                if child:IsA("RemoteEvent") or child:IsA("RemoteFunction") then
                    local name = child.Name
                    local isObfuscated = false
                    local reason = ""
                    if #name <= 3 and name:match("[%w]") then
                        isObfuscated = true
                        reason = "Very short name"
                    end
                    if name:match("^[a-zA-Z0-9]+$") and #name >= 8 then
                        local u, l, d = 0, 0, 0
                        for c in name:gmatch(".") do
                            if c:match("%u") then u += 1
                            elseif c:match("%l") then l += 1
                            elseif c:match("%d") then d += 1 end
                        end
                        if u > 2 and l > 2 and d > 1 then
                            isObfuscated = true
                            reason = "Random mixed-case"
                        end
                    end
                    if name:match("^[0-9a-fA-F]+$") and #name >= 6 then
                        isObfuscated = true
                        reason = "Hex-like encoding"
                    end
                    if isObfuscated then
                        local alreadyFound = false
                        for _, entry in ipairs(results.Remotes) do
                            if entry.Instance == child then alreadyFound = true break end
                        end
                        if not alreadyFound then
                            table.insert(results.Remotes, {
                                Instance = child, Name = child.Name, ClassName = child.ClassName,
                                Path = child:GetFullName(), Location = "Obfuscated",
                                MatchedKeyword = reason, ThreatLevel = "🟣 SUSPICIOUS",
                                ThreatColor = Theme.Purple, Category = AC_GetCategory(child),
                                Disabled = false, Hooked = false,
                            })
                            AC_TotalFound = AC_TotalFound + 1
                        end
                    end
                end
            end
        end)
    end
    local ReplicatedStorage = game:GetService("ReplicatedStorage")
    scanDesc(ReplicatedStorage)
    pcall(function() scanDesc(game:GetService("ReplicatedFirst")) end)
    pcall(function() scanDesc(workspace) end)
    pcall(function() scanDesc(LP.PlayerGui) end)
    pcall(function() scanDesc(game:GetService("Lighting")) end)
end

local function AC_FullScan(statusCallback)
    AC_IsScanning = true
    AC_ScanComplete = false
    AC_TotalFound = 0
    AC_TotalDisabled = 0
    AC_TotalHooked = 0
    AC_ScanResults = { Remotes = {}, Scripts = {}, Modules = {}, Misc = {}, Deleted = {} }

    local ReplicatedStorage = game:GetService("ReplicatedStorage")
    local scanLocations = {
        {ReplicatedStorage, "ReplicatedStorage"},
        {workspace, "Workspace"},
    }
    pcall(function() table.insert(scanLocations, {game:GetService("ReplicatedFirst"), "ReplicatedFirst"}) end)
    pcall(function() table.insert(scanLocations, {LP.PlayerGui, "PlayerGui"}) end)
    pcall(function() table.insert(scanLocations, {game:GetService("StarterPlayer"), "StarterPlayer"}) end)
    pcall(function() table.insert(scanLocations, {game:GetService("StarterGui"), "StarterGui"}) end)
    pcall(function() table.insert(scanLocations, {game:GetService("Lighting"), "Lighting"}) end)
    pcall(function() table.insert(scanLocations, {game:GetService("SoundService"), "SoundService"}) end)

    for _, loc in ipairs(scanLocations) do
        if statusCallback then statusCallback("🔍 Scanning " .. loc[2] .. "...") end
        task.wait()
        AC_ScanLocation(loc[1], loc[2], AC_ScanResults, statusCallback)
    end

    if statusCallback then statusCallback("🔍 Scanning for obfuscated remotes...") end
    task.wait()
    AC_ScanForObfuscated(AC_ScanResults, statusCallback)

    -- Check character
    pcall(function()
        local char = LP.Character
        if char then
            for _, child in pairs(char:GetDescendants()) do
                if child:IsA("LocalScript") or child:IsA("Script") then
                    local isAC, keyword = AC_IsACName(child.Name)
                    if isAC then
                        local tl, tc = AC_GetThreatLevel(child.Name)
                        table.insert(AC_ScanResults.Scripts, {
                            Instance = child, Name = child.Name, ClassName = child.ClassName,
                            Path = child:GetFullName(), Location = "Character",
                            MatchedKeyword = keyword, ThreatLevel = tl, ThreatColor = tc,
                            Category = AC_GetCategory(child), Disabled = false, Hooked = false,
                        })
                        AC_TotalFound += 1
                    end
                end
            end
        end
    end)

    local function threatSort(a, b)
        local order = {["🔴 HIGH"] = 1, ["🟣 SUSPICIOUS"] = 2, ["🟡 MEDIUM"] = 3, ["🟠 LOW"] = 4, ["⚪ UNKNOWN"] = 5}
        return (order[a.ThreatLevel] or 5) < (order[b.ThreatLevel] or 5)
    end
    table.sort(AC_ScanResults.Remotes, threatSort)
    table.sort(AC_ScanResults.Scripts, threatSort)
    table.sort(AC_ScanResults.Modules, threatSort)
    table.sort(AC_ScanResults.Misc, threatSort)

    AC_ScanComplete = true
    AC_IsScanning = false
    if statusCallback then
        statusCallback(string.format("✅ Scan Complete! Found %d suspicious items", AC_TotalFound))
    end
end

local function AC_DisableItem(entry)
    if not entry or not entry.Instance or not entry.Instance.Parent then return false end
    local ok = pcall(function()
        local item = entry.Instance
        if item:IsA("LocalScript") then
            item.Disabled = true
        elseif item:IsA("ModuleScript") then
            pcall(function() item.Disabled = true end)
        else
            item:Destroy()
        end
        entry.Disabled = true
        AC_TotalDisabled += 1
    end)
    return ok
end

local function AC_HookRemote(entry)
    if not entry or not entry.Instance or not entry.Instance.Parent then return false end
    AC_HookedRemotes[entry.Instance] = true
    entry.Hooked = true
    AC_TotalHooked += 1
    return true
end

local function AC_SetupNamecallHook()
    if AC_OriginalNamecallHook then return end
    pcall(function()
        if hookmetamethod then
            local oldNamecall
            oldNamecall = hookmetamethod(game, "__namecall", function(self, ...)
                local method = getnamecallmethod()
                if (method == "FireServer" or method == "InvokeServer") and AC_HookedRemotes[self] then
                    return nil
                end
                return oldNamecall(self, ...)
            end)
            AC_OriginalNamecallHook = oldNamecall
        end
    end)
end

local function AC_DisableHighThreats()
    local count = 0
    local allLists = {AC_ScanResults.Remotes, AC_ScanResults.Scripts, AC_ScanResults.Modules, AC_ScanResults.Misc}
    for _, list in ipairs(allLists) do
        for _, entry in ipairs(list) do
            if entry.ThreatLevel == "🔴 HIGH" and not entry.Disabled then
                if AC_DisableItem(entry) then count += 1 end
            end
        end
    end
    return count
end

local function AC_HookAllRemotes()
    local count = 0
    for _, entry in ipairs(AC_ScanResults.Remotes) do
        if not entry.Hooked and entry.Instance and entry.Instance.Parent then
            if AC_HookRemote(entry) then count += 1 end
        end
    end
    AC_SetupNamecallHook()
    return count
end

local function AC_NukeAll()
    local count = 0
    local allLists = {AC_ScanResults.Remotes, AC_ScanResults.Scripts, AC_ScanResults.Modules, AC_ScanResults.Misc}
    for _, list in ipairs(allLists) do
        for _, entry in ipairs(list) do
            if not entry.Disabled then
                if AC_DisableItem(entry) then count += 1 end
            end
        end
    end
    return count
end

-- ════════════════════════════════════════
-- // CHARACTER EVENTS
-- ════════════════════════════════════════
LP.CharacterAdded:Connect(function(char)
    if AutoReTP and LastTP then
        IsAtTarget = false
        task.spawn(function()
            local hrp = char:WaitForChild("HumanoidRootPart", 5)
            if hrp and AutoReTP and LastTP then DoInstantTP(LastTP) end
        end)
    end
    task.delay(1, RefreshGaps)
end)

StartMainTPLoop()

-- ════════════════════════════════════════
-- // GUI CREATION
-- ════════════════════════════════════════
local Gui = Instance.new("ScreenGui")
Gui.Name = "MuslimDevs"
Gui.ResetOnSpawn = false
Gui.ZIndexBehavior = Enum.ZIndexBehavior.Sibling
Gui.Parent = LP.PlayerGui

-- Floating Toggle Button
local FloatingBtn = Instance.new("TextButton")
FloatingBtn.Name = "FloatingToggle"
FloatingBtn.Size = UDim2.new(0, 50, 0, 50)
FloatingBtn.Position = UDim2.new(0, 10, 0.5, -25)
FloatingBtn.BackgroundColor3 = Theme.Primary
FloatingBtn.Text = "☪️"
FloatingBtn.TextSize = 24
FloatingBtn.Font = Enum.Font.GothamBold
FloatingBtn.TextColor3 = Theme.Text
FloatingBtn.BorderSizePixel = 0
FloatingBtn.AutoButtonColor = false
FloatingBtn.ZIndex = 100
FloatingBtn.Visible = false
FloatingBtn.Parent = Gui
Corner(FloatingBtn, 25)
Stroke(FloatingBtn, Theme.Gold, 2, 0)

task.spawn(function()
    local fbStroke = FloatingBtn:FindFirstChildOfClass("UIStroke")
    local hue = 0
    while fbStroke and fbStroke.Parent do
        hue = (hue + 0.01) % 1
        fbStroke.Color = Color3.fromHSV(hue, 0.7, 1)
        task.wait(0.03)
    end
end)

-- Floating button dragging
do
    local fbDragging, fbDragStart, fbStartPos
    FloatingBtn.InputBegan:Connect(function(input)
        if input.UserInputType == Enum.UserInputType.MouseButton1 or input.UserInputType == Enum.UserInputType.Touch then
            fbDragging = true
            fbDragStart = input.Position
            fbStartPos = FloatingBtn.Position
            input.Changed:Connect(function()
                if input.UserInputState == Enum.UserInputState.End then fbDragging = false end
            end)
        end
    end)
    FloatingBtn.InputChanged:Connect(function(input)
        if fbDragging and (input.UserInputType == Enum.UserInputType.MouseMovement or input.UserInputType == Enum.UserInputType.Touch) then
            local delta = input.Position - fbDragStart
            if delta.Magnitude > 5 then
                FloatingBtn.Position = UDim2.new(fbStartPos.X.Scale, fbStartPos.X.Offset + delta.X, fbStartPos.Y.Scale, fbStartPos.Y.Offset + delta.Y)
            end
        end
    end)
end

-- Main Frame
local Main = Instance.new("Frame")
Main.Name = "Main"
Main.Size = UDim2.new(0, MainWidth, 0, MainHeight)
Main.Position = UDim2.new(0.5, -MainWidth/2, 0.5, -MainHeight/2)
Main.BackgroundColor3 = Theme.Bg
Main.BorderSizePixel = 0
Main.ClipsDescendants = true
Main.Parent = Gui
Corner(Main, IsMobile and 14 or 18)
local mainStroke = Stroke(Main, Theme.PrimaryLight, IsMobile and 2 or 3, 0.1)

task.spawn(function()
    local hue = 0
    while mainStroke and mainStroke.Parent do
        hue = (hue + 0.005) % 1
        mainStroke.Color = Color3.fromHSV(hue, 0.8, 1)
        task.wait(0.02)
    end
end)

-- Header
local Header = Instance.new("Frame")
Header.Name = "Header"
Header.Size = UDim2.new(1, 0, 0, HeaderHeight)
Header.BackgroundColor3 = Theme.Bg
Header.BorderSizePixel = 0
Header.ZIndex = 10
Header.Parent = Main
Corner(Header, IsMobile and 14 or 18)

local headerGradient = Instance.new("UIGradient")
headerGradient.Color = ColorSequence.new({
    ColorSequenceKeypoint.new(0, Color3.fromRGB(20, 20, 40)),
    ColorSequenceKeypoint.new(0.5, Color3.fromRGB(40, 20, 60)),
    ColorSequenceKeypoint.new(1, Color3.fromRGB(20, 20, 40))
})
headerGradient.Parent = Header

task.spawn(function()
    local offset = 0
    while headerGradient and headerGradient.Parent do
        offset = (offset + 0.01) % 1
        headerGradient.Offset = Vector2.new(math.sin(offset * math.pi * 2) * 0.5, 0)
        task.wait(0.03)
    end
end)

local HeaderFix = Instance.new("Frame")
HeaderFix.Size = UDim2.new(1, 0, 0, 20)
HeaderFix.Position = UDim2.new(0, 0, 1, -20)
HeaderFix.BackgroundColor3 = Theme.Bg
HeaderFix.BorderSizePixel = 0
HeaderFix.ZIndex = 10
HeaderFix.Parent = Header

local LogoSize = IsMobile and 38 or 44
local Logo = Instance.new("TextLabel")
Logo.Size = UDim2.new(0, LogoSize, 0, LogoSize)
Logo.Position = UDim2.new(0, 8, 0.5, -LogoSize/2)
Logo.BackgroundColor3 = Color3.fromRGB(0, 0, 0)
Logo.BackgroundTransparency = 0.3
Logo.Text = "☪️"
Logo.TextSize = IsMobile and 22 or 26
Logo.Font = Enum.Font.GothamBold
Logo.TextColor3 = Theme.Text
Logo.ZIndex = 11
Logo.Parent = Header
Corner(Logo, 12)
Stroke(Logo, Theme.Gold, 2, 0)

task.spawn(function()
    local logoStroke = Logo:FindFirstChildOfClass("UIStroke")
    local hue = 0
    while logoStroke and logoStroke.Parent do
        hue = (hue + 0.01) % 1
        logoStroke.Color = Color3.fromHSV(hue, 0.7, 1)
        task.wait(0.03)
    end
end)

local titleX = LogoSize + 14
CreateRainbowText(Header, "MUSLIM DEVS", UDim2.new(0, IsMobile and 140 or 200, 0, 24), UDim2.new(0, titleX, 0, IsMobile and 6 or 8), IsMobile and 16 or 20)

local Sub = Instance.new("TextLabel")
Sub.Size = UDim2.new(0, IsMobile and 140 or 200, 0, 14)
Sub.Position = UDim2.new(0, titleX, 0, IsMobile and 30 or 36)
Sub.BackgroundTransparency = 1
Sub.Text = "Premium • " .. (IsMobile and "📱" or "💻") .. " " .. LP.Name
Sub.TextColor3 = Theme.Accent
Sub.TextSize = FS(9)
Sub.Font = Enum.Font.GothamBold
Sub.TextXAlignment = Enum.TextXAlignment.Left
Sub.ZIndex = 11
Sub.Parent = Header

task.spawn(function()
    local hue = 0
    while Sub and Sub.Parent do
        hue = (hue + 0.008) % 1
        Sub.TextColor3 = Color3.fromHSV(hue, 0.5, 1)
        task.wait(0.03)
    end
end)

local ctrlSize = IsMobile and 36 or 32

local CloseBtn = Instance.new("TextButton")
CloseBtn.Size = UDim2.new(0, ctrlSize, 0, ctrlSize)
CloseBtn.Position = UDim2.new(1, -(ctrlSize + 8), 0.5, -ctrlSize/2)
CloseBtn.BackgroundColor3 = Theme.Red
CloseBtn.Text = "✕"
CloseBtn.TextColor3 = Theme.Text
CloseBtn.TextSize = IsMobile and 16 or 14
CloseBtn.Font = Enum.Font.GothamBold
CloseBtn.BorderSizePixel = 0
CloseBtn.ZIndex = 12
CloseBtn.Parent = Header
Corner(CloseBtn, 8)
Ripple(CloseBtn)

local MinBtn = Instance.new("TextButton")
MinBtn.Size = UDim2.new(0, ctrlSize, 0, ctrlSize)
MinBtn.Position = UDim2.new(1, -(ctrlSize * 2 + 14), 0.5, -ctrlSize/2)
MinBtn.BackgroundColor3 = Theme.Yellow
MinBtn.Text = "—"
MinBtn.TextColor3 = Theme.Bg
MinBtn.TextSize = IsMobile and 18 or 16
MinBtn.Font = Enum.Font.GothamBold
MinBtn.BorderSizePixel = 0
MinBtn.ZIndex = 12
MinBtn.Parent = Header
Corner(MinBtn, 8)
Ripple(MinBtn)

-- Dragging
do
    local dragging, dragStart, startPos
    Header.InputBegan:Connect(function(input)
        if input.UserInputType == Enum.UserInputType.MouseButton1 or input.UserInputType == Enum.UserInputType.Touch then
            dragging = true
            dragStart = input.Position
            startPos = Main.Position
            input.Changed:Connect(function()
                if input.UserInputState == Enum.UserInputState.End then dragging = false end
            end)
        end
    end)
    UserInputService.InputChanged:Connect(function(input)
        if dragging and (input.UserInputType == Enum.UserInputType.MouseMovement or input.UserInputType == Enum.UserInputType.Touch) then
            local delta = input.Position - dragStart
            Main.Position = UDim2.new(startPos.X.Scale, startPos.X.Offset + delta.X, startPos.Y.Scale, startPos.Y.Offset + delta.Y)
        end
    end)
end

-- Tab System
local TabBar = Instance.new("Frame")
TabBar.Size = UDim2.new(0, TabWidth, 1, -(HeaderHeight + 6))
TabBar.Position = UDim2.new(0, 4, 0, HeaderHeight + 2)
TabBar.BackgroundColor3 = Theme.BgLight
TabBar.BorderSizePixel = 0
TabBar.ZIndex = 5
TabBar.ClipsDescendants = true
TabBar.Parent = Main
Corner(TabBar, 10)

local TabScroll = Instance.new("ScrollingFrame")
TabScroll.Size = UDim2.new(1, 0, 1, 0)
TabScroll.BackgroundTransparency = 1
TabScroll.ScrollBarThickness = 0
TabScroll.CanvasSize = UDim2.new(0, 0, 0, 0)
TabScroll.AutomaticCanvasSize = Enum.AutomaticSize.Y
TabScroll.ScrollingDirection = Enum.ScrollingDirection.Y
TabScroll.Parent = TabBar

local TabLayout = Instance.new("UIListLayout")
TabLayout.FillDirection = Enum.FillDirection.Vertical
TabLayout.HorizontalAlignment = Enum.HorizontalAlignment.Center
TabLayout.Padding = UDim.new(0, IsMobile and 3 or 4)
TabLayout.Parent = TabScroll

local TabPad = Instance.new("UIPadding")
TabPad.PaddingTop = UDim.new(0, 4)
TabPad.PaddingLeft = UDim.new(0, 3)
TabPad.PaddingRight = UDim.new(0, 3)
TabPad.PaddingBottom = UDim.new(0, 4)
TabPad.Parent = TabScroll

local contentX = TabWidth + 10
local Content = Instance.new("Frame")
Content.Size = UDim2.new(1, -(contentX + 4), 1, -(HeaderHeight + 6))
Content.Position = UDim2.new(0, contentX, 0, HeaderHeight + 2)
Content.BackgroundTransparency = 1
Content.ClipsDescendants = true
Content.ZIndex = 3
Content.Parent = Main

local Pages = {}
local TabBtns = {}
local ActiveTab = nil

-- ════════════════════════════════════════
-- // TAB DEFINITIONS
-- ════════════════════════════════════════
local TabDefs = {
    {Name = "Gaps", Icon = "📍"},
    {Name = "ESP", Icon = "👁️"},
    {Name = "Coins", Icon = "🪙"},
    {Name = "Money", Icon = "💰"},
    {Name = "Auto", Icon = "🔄"},
    {Name = "Grab", Icon = "🫳"},
    {Name = "Walls", Icon = "🧱"},
    {Name = "Info", Icon = "💬"},
}

for _, def in ipairs(TabDefs) do
    local page = Instance.new("ScrollingFrame")
    page.Name = def.Name
    page.Size = UDim2.new(1, 0, 1, 0)
    page.BackgroundTransparency = 1
    page.BorderSizePixel = 0
    page.ScrollBarThickness = IsMobile and 4 or 3
    page.ScrollBarImageColor3 = Theme.Primary
    page.CanvasSize = UDim2.new(0, 0, 0, 0)
    page.AutomaticCanvasSize = Enum.AutomaticSize.Y
    page.Visible = false
    page.ZIndex = 3
    page.Parent = Content

    local pl = Instance.new("UIListLayout")
    pl.Padding = UDim.new(0, IsMobile and 5 or 6)
    pl.Parent = page

    local pp = Instance.new("UIPadding")
    pp.PaddingTop = UDim.new(0, 4)
    pp.PaddingBottom = UDim.new(0, 12)
    pp.PaddingLeft = UDim.new(0, 3)
    pp.PaddingRight = UDim.new(0, 3)
    pp.Parent = page

    Pages[def.Name] = page
end

local function SwitchTab(name)
    if ActiveTab == name then return end
    ActiveTab = name
    for pName, page in pairs(Pages) do page.Visible = (pName == name) end
    for bName, btn in pairs(TabBtns) do
        Tween(btn, {BackgroundColor3 = (bName == name) and Theme.Primary or Theme.BgLighter}, 0.2)
    end
end

for i, def in ipairs(TabDefs) do
    local btn = Instance.new("TextButton")
    btn.Name = def.Name
    btn.Size = UDim2.new(1, 0, 0, TabBtnH)
    btn.BackgroundColor3 = (i == 1) and Theme.Primary or Theme.BgLighter
    btn.BorderSizePixel = 0
    btn.Text = ""
    btn.AutoButtonColor = false
    btn.LayoutOrder = i
    btn.ZIndex = 6
    btn.Parent = TabScroll
    Corner(btn, 8)

    local icon = Instance.new("TextLabel")
    icon.Size = UDim2.new(1, 0, 0, 20)
    icon.Position = UDim2.new(0, 0, 0, IsMobile and 6 or 4)
    icon.BackgroundTransparency = 1
    icon.Text = def.Icon
    icon.TextSize = TabIconSize
    icon.Font = Enum.Font.GothamBold
    icon.TextColor3 = Theme.Text
    icon.ZIndex = 6
    icon.Parent = btn

    local lbl = Instance.new("TextLabel")
    lbl.Size = UDim2.new(1, 0, 0, 14)
    lbl.Position = UDim2.new(0, 0, 0, IsMobile and 26 or 24)
    lbl.BackgroundTransparency = 1
    lbl.Text = def.Name
    lbl.TextColor3 = Theme.TextDim
    lbl.TextSize = TabLabelSize
    lbl.Font = Enum.Font.GothamBold
    lbl.ZIndex = 6
    lbl.Parent = btn

    btn.MouseButton1Click:Connect(function() SwitchTab(def.Name) end)
    Ripple(btn)
    TabBtns[def.Name] = btn
end

-- ════════════════════════════════════════
-- // UI BUILDERS
-- ════════════════════════════════════════
local function MakeCard(parent, height)
    local c = Instance.new("Frame")
    c.Size = UDim2.new(1, 0, 0, height)
    c.BackgroundColor3 = Theme.BgLight
    c.BorderSizePixel = 0
    c.Parent = parent
    Corner(c, IsMobile and 8 or 10)
    return c
end

local function MakeLabel(parent, text, color, size)
    local l = Instance.new("TextLabel")
    l.Size = UDim2.new(1, 0, 0, IsMobile and 22 or 20)
    l.BackgroundTransparency = 1
    l.Text = "  " .. text
    l.TextColor3 = color or Theme.Accent
    l.TextSize = FS(size or 12)
    l.Font = Enum.Font.GothamBlack
    l.TextXAlignment = Enum.TextXAlignment.Left
    l.Parent = parent
    return l
end

local function MakeRainbowLabel(parent, text, size)
    local container = Instance.new("Frame")
    container.Size = UDim2.new(1, 0, 0, IsMobile and 22 or 20)
    container.BackgroundTransparency = 1
    container.Parent = parent
    CreateRainbowText(container, "  " .. text, UDim2.new(1, 0, 1, 0), UDim2.new(0, 0, 0, 0), FS(size or 12))
    return container
end

local function MakeToggle(parent, text, default, callback)
    local card = MakeCard(parent, ToggleH)
    local lbl = Instance.new("TextLabel")
    lbl.Size = UDim2.new(0.65, 0, 1, 0)
    lbl.Position = UDim2.new(0, CardPad, 0, 0)
    lbl.BackgroundTransparency = 1
    lbl.Text = text
    lbl.TextColor3 = Theme.Text
    lbl.TextSize = FS(IsMobile and 10 or 11)
    lbl.Font = Enum.Font.GothamBold
    lbl.TextXAlignment = Enum.TextXAlignment.Left
    lbl.TextWrapped = true
    lbl.Parent = card

    local toggleW = IsMobile and 50 or 44
    local toggleHLocal = IsMobile and 26 or 22
    local bg = Instance.new("Frame")
    bg.Size = UDim2.new(0, toggleW, 0, toggleHLocal)
    bg.Position = UDim2.new(1, -(toggleW + 8), 0.5, -toggleHLocal/2)
    bg.BackgroundColor3 = default and Theme.Green or Theme.BgLighter
    bg.BorderSizePixel = 0
    bg.Parent = card
    Corner(bg, toggleHLocal/2)

    local knobSize = toggleHLocal - 4
    local knob = Instance.new("Frame")
    knob.Size = UDim2.new(0, knobSize, 0, knobSize)
    knob.Position = default and UDim2.new(1, -(knobSize + 2), 0.5, -knobSize/2) or UDim2.new(0, 2, 0.5, -knobSize/2)
    knob.BackgroundColor3 = Theme.Text
    knob.BorderSizePixel = 0
    knob.Parent = bg
    Corner(knob, knobSize/2)

    local isOn = default or false
    local btn = Instance.new("TextButton")
    btn.Size = UDim2.new(1, 0, 1, 0)
    btn.BackgroundTransparency = 1
    btn.Text = ""
    btn.Parent = card

    btn.MouseButton1Click:Connect(function()
        isOn = not isOn
        Tween(knob, {Position = isOn and UDim2.new(1, -(knobSize + 2), 0.5, -knobSize/2) or UDim2.new(0, 2, 0.5, -knobSize/2)}, 0.2)
        Tween(bg, {BackgroundColor3 = isOn and Theme.Green or Theme.BgLighter}, 0.2)
        if callback then callback(isOn) end
    end)
    return card, bg, knob, isOn
end

local function MakeButton(parent, text, color, callback)
    local btn = Instance.new("TextButton")
    btn.Size = UDim2.new(1, 0, 0, BtnHeight)
    btn.BackgroundColor3 = color or Theme.BgLight
    btn.Text = text
    btn.TextColor3 = Theme.Text
    btn.TextSize = FS(IsMobile and 11 or 11)
    btn.Font = Enum.Font.GothamBlack
    btn.BorderSizePixel = 0
    btn.AutoButtonColor = false
    btn.Parent = parent
    Corner(btn, 8)
    Ripple(btn)
    Hover(btn, color or Theme.BgLight, Color3.new(
        math.min((color or Theme.BgLight).R + 0.08, 1),
        math.min((color or Theme.BgLight).G + 0.08, 1),
        math.min((color or Theme.BgLight).B + 0.08, 1)
    ))
    btn.MouseButton1Click:Connect(function() if callback then callback(btn) end end)
    return btn
end

local function MakeSlider(parent, text, min, max, default, suffix, callback)
    local sliderH = IsMobile and 56 or 50
    local card = MakeCard(parent, sliderH)

    local lbl = Instance.new("TextLabel")
    lbl.Size = UDim2.new(0.55, 0, 0, 18)
    lbl.Position = UDim2.new(0, CardPad, 0, 3)
    lbl.BackgroundTransparency = 1
    lbl.Text = text
    lbl.TextColor3 = Theme.Text
    lbl.TextSize = FS(10)
    lbl.Font = Enum.Font.GothamBold
    lbl.TextXAlignment = Enum.TextXAlignment.Left
    lbl.Parent = card

    local valLbl = Instance.new("TextLabel")
    valLbl.Size = UDim2.new(0.4, -CardPad, 0, 18)
    valLbl.Position = UDim2.new(0.6, 0, 0, 3)
    valLbl.BackgroundTransparency = 1
    valLbl.Text = tostring(default) .. (suffix or "")
    valLbl.TextColor3 = Theme.Accent
    valLbl.TextSize = FS(10)
    valLbl.Font = Enum.Font.GothamBlack
    valLbl.TextXAlignment = Enum.TextXAlignment.Right
    valLbl.Parent = card

    local trackH = IsMobile and 14 or 8
    local trackY = IsMobile and 30 or 32
    local track = Instance.new("Frame")
    track.Size = UDim2.new(1, -CardPad * 2, 0, trackH)
    track.Position = UDim2.new(0, CardPad, 0, trackY)
    track.BackgroundColor3 = Theme.BgLighter
    track.BorderSizePixel = 0
    track.Parent = card
    Corner(track, trackH/2)

    local fill = Instance.new("Frame")
    fill.Size = UDim2.new((default - min) / (max - min), 0, 1, 0)
    fill.BackgroundColor3 = Theme.Primary
    fill.BorderSizePixel = 0
    fill.Parent = track
    Corner(fill, trackH/2)

    local thumb = Instance.new("Frame")
    thumb.Size = UDim2.new(0, IsMobile and 22 or 14, 0, IsMobile and 22 or 14)
    thumb.AnchorPoint = Vector2.new(0.5, 0.5)
    thumb.Position = UDim2.new((default - min) / (max - min), 0, 0.5, 0)
    thumb.BackgroundColor3 = Theme.Text
    thumb.BorderSizePixel = 0
    thumb.ZIndex = 4
    thumb.Parent = track
    Corner(thumb, IsMobile and 11 or 7)

    local hitbox = Instance.new("TextButton")
    hitbox.Size = UDim2.new(1, 10, 1, IsMobile and 20 or 12)
    hitbox.Position = UDim2.new(0, -5, 0, IsMobile and -10 or -6)
    hitbox.BackgroundTransparency = 1
    hitbox.Text = ""
    hitbox.ZIndex = 5
    hitbox.Parent = track

    local sliding = false
    local currentInputPos = nil

    hitbox.InputBegan:Connect(function(input)
        if input.UserInputType == Enum.UserInputType.MouseButton1 or input.UserInputType == Enum.UserInputType.Touch then
            sliding = true
            currentInputPos = input.Position
        end
    end)

    hitbox.InputChanged:Connect(function(input)
        if sliding and (input.UserInputType == Enum.UserInputType.MouseMovement or input.UserInputType == Enum.UserInputType.Touch) then
            currentInputPos = input.Position
        end
    end)

    UserInputService.InputChanged:Connect(function(input)
        if sliding and (input.UserInputType == Enum.UserInputType.MouseMovement or input.UserInputType == Enum.UserInputType.Touch) then
            currentInputPos = input.Position
        end
    end)

    UserInputService.InputEnded:Connect(function(input)
        if input.UserInputType == Enum.UserInputType.MouseButton1 or input.UserInputType == Enum.UserInputType.Touch then
            sliding = false
        end
    end)

    local conn
    conn = RunService.RenderStepped:Connect(function()
        if not card or not card.Parent then conn:Disconnect() return end
        if sliding and currentInputPos then
            local rel = math.clamp((currentInputPos.X - track.AbsolutePosition.X) / track.AbsoluteSize.X, 0, 1)
            local val = math.floor(min + (max - min) * rel)
            fill.Size = UDim2.new(rel, 0, 1, 0)
            thumb.Position = UDim2.new(rel, 0, 0.5, 0)
            valLbl.Text = tostring(val) .. (suffix or "")
            if callback then callback(val) end
        end
    end)
    return card
end

-- ════════════════════════════════════════════════════════════════
-- // ESP PAGE CONTENT
-- ════════════════════════════════════════════════════════════════
local ESP_Page = Pages["ESP"]
local ESPFolderFrames = {}
local ESPStatusLabel = nil

MakeRainbowLabel(ESP_Page, "👁️ BRAINROT ESP SYSTEM", 12)

-- ESP Info Card
local espInfoCard = MakeCard(ESP_Page, IsMobile and 50 or 55)
local espInfoText = Instance.new("TextLabel")
espInfoText.Size = UDim2.new(1, -12, 1, 0)
espInfoText.Position = UDim2.new(0, 6, 0, 0)
espInfoText.BackgroundTransparency = 1
espInfoText.Text = "📍 Draws ESP on brainrots!\n📦 Box | 📝 Name | 📏 Distance | 📍 Tracer"
espInfoText.TextColor3 = Theme.TextDim
espInfoText.TextSize = FS(9)
espInfoText.Font = Enum.Font.Gotham
espInfoText.TextWrapped = true
espInfoText.TextXAlignment = Enum.TextXAlignment.Left
espInfoText.Parent = espInfoCard

Instance.new("Frame", ESP_Page).Size = UDim2.new(1, 0, 0, 4)

-- Master ESP Toggle
MakeToggle(ESP_Page, "👁️ Enable ESP System", false, function(state)
    ESPConfig.Enabled = state
    if not state then
        for _, data in pairs(ESPObjects) do
            HideESP(data.esp)
        end
    end
end)

Instance.new("Frame", ESP_Page).Size = UDim2.new(1, 0, 0, 4)
MakeLabel(ESP_Page, "⚙️ ESP OPTIONS", Theme.ESP, 10)

MakeToggle(ESP_Page, "📦 Show Box", ESPConfig.ShowBox, function(state)
    ESPConfig.ShowBox = state
end)

MakeToggle(ESP_Page, "📝 Show Name", ESPConfig.ShowName, function(state)
    ESPConfig.ShowName = state
end)

MakeToggle(ESP_Page, "🏷️ Show Rarity", ESPConfig.ShowRarity, function(state)
    ESPConfig.ShowRarity = state
end)

MakeToggle(ESP_Page, "📏 Show Distance", ESPConfig.ShowDistance, function(state)
    ESPConfig.ShowDistance = state
end)

MakeToggle(ESP_Page, "📍 Show Tracer", ESPConfig.ShowTracer, function(state)
    ESPConfig.ShowTracer = state
end)

Instance.new("Frame", ESP_Page).Size = UDim2.new(1, 0, 0, 4)

MakeSlider(ESP_Page, "📏 Max Distance", 100, 5000, 2000, " studs", function(val)
    ESPConfig.MaxDistance = val
end)

MakeSlider(ESP_Page, "📝 Text Size", 10, 24, 14, "px", function(val)
    ESPConfig.TextSize = val
end)

Instance.new("Frame", ESP_Page).Size = UDim2.new(1, 0, 0, 4)
MakeLabel(ESP_Page, "📁 SELECT RARITIES", Theme.ESP, 10)

-- ESP Status
local espStatusCard = MakeCard(ESP_Page, IsMobile and 34 or 30)
ESPStatusLabel = Instance.new("TextLabel")
ESPStatusLabel.Size = UDim2.new(1, -12, 1, 0)
ESPStatusLabel.Position = UDim2.new(0, 6, 0, 0)
ESPStatusLabel.BackgroundTransparency = 1
ESPStatusLabel.Text = "❌ No selection - ESP disabled"
ESPStatusLabel.TextColor3 = Theme.Red
ESPStatusLabel.TextSize = FS(10)
ESPStatusLabel.Font = Enum.Font.GothamBold
ESPStatusLabel.TextXAlignment = Enum.TextXAlignment.Left
ESPStatusLabel.Parent = espStatusCard

local function UpdateESPStatus()
    local folderCount = 0
    local brainrotCount = 0
    
    for _, v in pairs(SelectedFolders) do 
        if v then folderCount = folderCount + 1 end 
    end
    for _, v in pairs(SelectedBrainrots) do 
        if v then brainrotCount = brainrotCount + 1 end 
    end

    if ESPStatusLabel and ESPStatusLabel.Parent then
        if folderCount == 0 and brainrotCount == 0 then
            ESPStatusLabel.Text = "❌ No selection - ESP disabled"
            ESPStatusLabel.TextColor3 = Theme.Red
        else
            ESPStatusLabel.Text = string.format("✅ Active: %d folders, %d individual", folderCount, brainrotCount)
            ESPStatusLabel.TextColor3 = Theme.Green
        end
    end
end

-- Toolbar for ESP
local espToolbar = Instance.new("Frame")
espToolbar.Size = UDim2.new(1, 0, 0, BtnHeight)
espToolbar.BackgroundTransparency = 1
espToolbar.Parent = ESP_Page

local espSelectAllBtn = Instance.new("TextButton")
espSelectAllBtn.Size = UDim2.new(0.32, -2, 1, 0)
espSelectAllBtn.Position = UDim2.new(0, 0, 0, 0)
espSelectAllBtn.BackgroundColor3 = Theme.Green
espSelectAllBtn.Text = "✅ All"
espSelectAllBtn.TextColor3 = Theme.Text
espSelectAllBtn.TextSize = FS(10)
espSelectAllBtn.Font = Enum.Font.GothamBold
espSelectAllBtn.BorderSizePixel = 0
espSelectAllBtn.Parent = espToolbar
Corner(espSelectAllBtn, 6)
Ripple(espSelectAllBtn)

local espDeselectAllBtn = Instance.new("TextButton")
espDeselectAllBtn.Size = UDim2.new(0.32, -2, 1, 0)
espDeselectAllBtn.Position = UDim2.new(0.33, 0, 0, 0)
espDeselectAllBtn.BackgroundColor3 = Theme.Red
espDeselectAllBtn.Text = "❌ None"
espDeselectAllBtn.TextColor3 = Theme.Text
espDeselectAllBtn.TextSize = FS(10)
espDeselectAllBtn.Font = Enum.Font.GothamBold
espDeselectAllBtn.BorderSizePixel = 0
espDeselectAllBtn.Parent = espToolbar
Corner(espDeselectAllBtn, 6)
Ripple(espDeselectAllBtn)

local espRefreshBtn = Instance.new("TextButton")
espRefreshBtn.Size = UDim2.new(0.32, -2, 1, 0)
espRefreshBtn.Position = UDim2.new(0.66, 2, 0, 0)
espRefreshBtn.BackgroundColor3 = Theme.Blue
espRefreshBtn.Text = "🔄 Refresh"
espRefreshBtn.TextColor3 = Theme.Text
espRefreshBtn.TextSize = FS(10)
espRefreshBtn.Font = Enum.Font.GothamBold
espRefreshBtn.BorderSizePixel = 0
espRefreshBtn.Parent = espToolbar
Corner(espRefreshBtn, 6)
Ripple(espRefreshBtn)

-- ESP Folder List Container
local espScrollFrame = MakeCard(ESP_Page, IsMobile and 160 or 180)
espScrollFrame.ClipsDescendants = true

local espInnerScroll = Instance.new("ScrollingFrame")
espInnerScroll.Size = UDim2.new(1, -4, 1, -4)
espInnerScroll.Position = UDim2.new(0, 2, 0, 2)
espInnerScroll.BackgroundTransparency = 1
espInnerScroll.ScrollBarThickness = IsMobile and 5 or 3
espInnerScroll.ScrollBarImageColor3 = Theme.ESP
espInnerScroll.CanvasSize = UDim2.new(0, 0, 0, 0)
espInnerScroll.AutomaticCanvasSize = Enum.AutomaticSize.Y
espInnerScroll.Parent = espScrollFrame

local espListLayout = Instance.new("UIListLayout")
espListLayout.Padding = UDim.new(0, 4)
espListLayout.Parent = espInnerScroll

local espPadding = Instance.new("UIPadding")
espPadding.PaddingTop = UDim.new(0, 2)
espPadding.PaddingBottom = UDim.new(0, 2)
espPadding.PaddingLeft = UDim.new(0, 2)
espPadding.PaddingRight = UDim.new(0, 2)
espPadding.Parent = espInnerScroll

local function RefreshESPUI()
    for _, frame in pairs(ESPFolderFrames) do
        pcall(function() frame:Destroy() end)
    end
    ESPFolderFrames = {}

    local structure = ScanBrainrotFolders()

    if not next(structure) then
        local emptyLabel = Instance.new("TextLabel")
        emptyLabel.Size = UDim2.new(1, 0, 0, 40)
        emptyLabel.BackgroundTransparency = 1
        emptyLabel.Text = "No brainrots found!\nWait for them to spawn..."
        emptyLabel.TextColor3 = Theme.TextMuted
        emptyLabel.TextSize = FS(10)
        emptyLabel.Font = Enum.Font.Gotham
        emptyLabel.Parent = espInnerScroll
        table.insert(ESPFolderFrames, emptyLabel)
        UpdateESPStatus()
        return
    end

    for folderName, brainrotList in pairs(structure) do
        local color = GetRarityColor(folderName)

        local folderBtn = Instance.new("TextButton")
        folderBtn.Name = folderName
        folderBtn.Size = UDim2.new(1, -4, 0, IsMobile and 36 or 32)
        folderBtn.BackgroundColor3 = SelectedFolders[folderName] and Color3.fromRGB(color.R * 255 * 0.3, color.G * 255 * 0.3, color.B * 255 * 0.3) or Theme.BgLighter
        folderBtn.BorderSizePixel = 0
        folderBtn.Text = ""
        folderBtn.AutoButtonColor = false
        folderBtn.Parent = espInnerScroll
        Corner(folderBtn, 6)
        
        local fStroke = Stroke(folderBtn, color, 1, 0.5)

        local colorBar = Instance.new("Frame")
        colorBar.Size = UDim2.new(0, 4, 1, -8)
        colorBar.Position = UDim2.new(0, 4, 0, 4)
        colorBar.BackgroundColor3 = color
        colorBar.BorderSizePixel = 0
        colorBar.Parent = folderBtn
        Corner(colorBar, 2)

        local folderLabel = Instance.new("TextLabel")
        folderLabel.Size = UDim2.new(1, -70, 1, 0)
        folderLabel.Position = UDim2.new(0, 14, 0, 0)
        folderLabel.BackgroundTransparency = 1
        folderLabel.Text = "📁 " .. folderName .. " (" .. #brainrotList .. ")"
        folderLabel.TextColor3 = color
        folderLabel.TextSize = FS(10)
        folderLabel.Font = Enum.Font.GothamBold
        folderLabel.TextXAlignment = Enum.TextXAlignment.Left
        folderLabel.Parent = folderBtn

        local checkLabel = Instance.new("TextLabel")
        checkLabel.Size = UDim2.new(0, 30, 1, 0)
        checkLabel.Position = UDim2.new(1, -35, 0, 0)
        checkLabel.BackgroundTransparency = 1
        checkLabel.Text = SelectedFolders[folderName] and "✅" or "⬜"
        checkLabel.TextSize = FS(14)
        checkLabel.Parent = folderBtn

        folderBtn.MouseButton1Click:Connect(function()
            SelectedFolders[folderName] = not SelectedFolders[folderName]
            checkLabel.Text = SelectedFolders[folderName] and "✅" or "⬜"
            folderBtn.BackgroundColor3 = SelectedFolders[folderName] and Color3.fromRGB(color.R * 255 * 0.3, color.G * 255 * 0.3, color.B * 255 * 0.3) or Theme.BgLighter
            UpdateESPStatus()
        end)

        table.insert(ESPFolderFrames, folderBtn)
    end

    UpdateESPStatus()
end

espSelectAllBtn.MouseButton1Click:Connect(function()
    local structure = ScanBrainrotFolders()
    for folderName, _ in pairs(structure) do
        SelectedFolders[folderName] = true
    end
    RefreshESPUI()
end)

espDeselectAllBtn.MouseButton1Click:Connect(function()
    ClearAllESPSelections()
    RefreshESPUI()
end)

espRefreshBtn.MouseButton1Click:Connect(function()
    espRefreshBtn.Text = "⏳..."
    task.wait(0.2)
    RefreshESPUI()
    espRefreshBtn.Text = "✅ Done!"
    task.delay(1, function() espRefreshBtn.Text = "🔄 Refresh" end)
end)

-- ════════════════════════════════════════════════════════════════
-- // GAPS PAGE CONTENT
-- ════════════════════════════════════════════════════════════════
local GP = Pages["Gaps"]

MakeRainbowLabel(GP, "📍 INSTANT MUD TELEPORTER", 12)

local gapStatsCard = MakeCard(GP, IsMobile and 48 or 52)
local GapCountLabel = Instance.new("TextLabel")
GapCountLabel.Size = UDim2.new(0.5, 0, 0, 20)
GapCountLabel.Position = UDim2.new(0, 8, 0, 4)
GapCountLabel.BackgroundTransparency = 1
GapCountLabel.Text = "🕳️ Gaps: 0"
GapCountLabel.TextColor3 = Theme.Green
GapCountLabel.TextSize = FS(11)
GapCountLabel.Font = Enum.Font.GothamBold
GapCountLabel.TextXAlignment = Enum.TextXAlignment.Left
GapCountLabel.Parent = gapStatsCard

local LastTPLabel = Instance.new("TextLabel")
LastTPLabel.Size = UDim2.new(0.5, -8, 0, 20)
LastTPLabel.Position = UDim2.new(0.5, 0, 0, 4)
LastTPLabel.BackgroundTransparency = 1
LastTPLabel.Text = "🎯 Last: None"
LastTPLabel.TextColor3 = Theme.TextDim
LastTPLabel.TextSize = FS(9)
LastTPLabel.Font = Enum.Font.GothamBold
LastTPLabel.TextXAlignment = Enum.TextXAlignment.Right
LastTPLabel.Parent = gapStatsCard

local MapDetectLabel = Instance.new("TextLabel")
MapDetectLabel.Size = UDim2.new(1, -16, 0, 16)
MapDetectLabel.Position = UDim2.new(0, 8, 0, 26)
MapDetectLabel.BackgroundTransparency = 1
MapDetectLabel.Text = "🗺️ Map: Detecting..."
MapDetectLabel.TextColor3 = Theme.TextMuted
MapDetectLabel.TextSize = FS(9)
MapDetectLabel.Font = Enum.Font.Gotham
MapDetectLabel.TextXAlignment = Enum.TextXAlignment.Left
MapDetectLabel.Parent = gapStatsCard

MakeToggle(GP, "🔁 Re-TP (INSTANT)", false, function(state)
    AutoReTP = state
    IsAtTarget = false
end)

MakeButton(GP, "💾 SAVE POSITION", Theme.Blue, function(btn)
    local hrp = GetHRP()
    if hrp then
        LastTP = hrp.CFrame
        IsAtTarget = true
        btn.Text = "✅ SAVED!"
        LastTPLabel.Text = "🎯 Custom Pos"
        task.delay(1.5, function() btn.Text = "💾 SAVE POSITION" end)
    end
end)

MakeButton(GP, "🗑️ CLEAR TARGET", Theme.Red, function(btn)
    LastTP = nil
    LastTPIndex = nil
    IsAtTarget = false
    btn.Text = "✅ CLEARED!"
    LastTPLabel.Text = "🎯 Last: None"
    task.delay(1, function() btn.Text = "🗑️ CLEAR TARGET" end)
end)

local gapFrameH = IsMobile and 140 or 120
local GapFrame = MakeCard(GP, gapFrameH)
GapFrame.ClipsDescendants = true

local GapScroll = Instance.new("ScrollingFrame")
GapScroll.Size = UDim2.new(1, -4, 1, -4)
GapScroll.Position = UDim2.new(0, 2, 0, 2)
GapScroll.BackgroundTransparency = 1
GapScroll.ScrollBarThickness = IsMobile and 5 or 3
GapScroll.ScrollBarImageColor3 = Theme.Primary
GapScroll.CanvasSize = UDim2.new(0, 0, 0, 0)
GapScroll.AutomaticCanvasSize = Enum.AutomaticSize.Y
GapScroll.Parent = GapFrame

Instance.new("UIListLayout", GapScroll).Padding = UDim.new(0, IsMobile and 4 or 3)

local GapButtons = {}

local function PopulateGaps()
    for _, b in ipairs(GapButtons) do if b and b.Parent then b:Destroy() end end
    GapButtons = {}
    for i, gapData in ipairs(AllGaps) do
        local row = Instance.new("Frame")
        row.Size = UDim2.new(1, -4, 0, GapRowH)
        row.BackgroundColor3 = Theme.BgLighter
        row.BorderSizePixel = 0
        row.LayoutOrder = i
        row.Parent = GapScroll
        Corner(row, 6)

        local mapColor = Theme.Primary
        local mapEmoji = "🌋"
        if gapData.Map == "DoomMap" then mapColor = Theme.Doom mapEmoji = "🔥"
        elseif gapData.Map == "ValentinesMap" then mapColor = Theme.Pink mapEmoji = "💝" end

        local numSize = IsMobile and 26 or 24
        local num = Instance.new("TextLabel")
        num.Size = UDim2.new(0, numSize, 0, numSize - 4)
        num.Position = UDim2.new(0, 3, 0.5, -(numSize-4)/2)
        num.BackgroundColor3 = mapColor
        num.Text = tostring(i)
        num.TextColor3 = Theme.Text
        num.TextSize = FS(9)
        num.Font = Enum.Font.GothamBlack
        num.Parent = row
        Corner(num, 4)

        local name = Instance.new("TextLabel")
        name.Size = UDim2.new(0, IsMobile and 50 or 60, 1, 0)
        name.Position = UDim2.new(0, numSize + 6, 0, 0)
        name.BackgroundTransparency = 1
        name.Text = gapData.Name
        name.TextColor3 = Theme.Accent
        name.TextSize = FS(9)
        name.Font = Enum.Font.GothamBold
        name.TextXAlignment = Enum.TextXAlignment.Left
        name.TextTruncate = Enum.TextTruncate.AtEnd
        name.Parent = row

        local goBtnW = IsMobile and 44 or 36
        local goBtnH = IsMobile and 28 or 22
        local tpBtn = Instance.new("TextButton")
        tpBtn.Size = UDim2.new(0, goBtnW, 0, goBtnH)
        tpBtn.Position = UDim2.new(1, -(goBtnW + 4), 0.5, -goBtnH/2)
        tpBtn.BackgroundColor3 = Theme.Green
        tpBtn.Text = "GO"
        tpBtn.TextColor3 = Theme.Bg
        tpBtn.TextSize = FS(IsMobile and 11 or 10)
        tpBtn.Font = Enum.Font.GothamBlack
        tpBtn.BorderSizePixel = 0
        tpBtn.AutoButtonColor = false
        tpBtn.Parent = row
        Corner(tpBtn, 4)
        Ripple(tpBtn)

        tpBtn.MouseButton1Click:Connect(function()
            TPToMud(i)
            LastTPLabel.Text = "🎯 " .. mapEmoji .. gapData.Name
        end)

        table.insert(GapButtons, row)
    end

    GapCountLabel.Text = "🕳️ Gaps: " .. #AllGaps
    local defaultCount, doomCount, valCount = 0, 0, 0
    for _, gap in ipairs(AllGaps) do
        if gap.Map == "DoomMap" then doomCount += 1
        elseif gap.Map == "ValentinesMap" then valCount += 1
        else defaultCount += 1 end
    end
    MapDetectLabel.Text = string.format("🗺️ %s | 🌋%d 🔥%d 💝%d", CurrentMap, defaultCount, doomCount, valCount)
    MapDetectLabel.TextColor3 = CurrentMap == "DoomMap" and Theme.Doom or CurrentMap == "ValentinesMap" and Theme.Pink or Theme.Blue
end

MakeButton(GP, "🔄 REFRESH GAPS", Theme.Primary, function(btn)
    btn.Text = "⏳ SCANNING..."
    task.wait(0.2)
    RefreshGaps()
    PopulateGaps()
    btn.Text = "✅ FOUND " .. #AllGaps .. "!"
    task.delay(1.5, function() btn.Text = "🔄 REFRESH GAPS" end)
end)

-- ════════════════════════════════════════════════════════════════
-- // SAFE TP TO SPAWN SECTION (NEW!)
-- ════════════════════════════════════════════════════════════════
Instance.new("Frame", GP).Size = UDim2.new(1, 0, 0, 8)
MakeRainbowLabel(GP, "🛡️ SAFE UNDERGROUND TP", 12)

-- Info card explaining the feature
local safeInfoCard = MakeCard(GP, IsMobile and 52 or 48)
local safeInfoText = Instance.new("TextLabel")
safeInfoText.Size = UDim2.new(1, -12, 1, 0)
safeInfoText.Position = UDim2.new(0, 6, 0, 0)
safeInfoText.BackgroundTransparency = 1
safeInfoText.Text = "🌊 Avoids surface waves by traveling UNDERGROUND!\n⬇️ Dive → ➡️ Move Below → ⬆️ Rise at Spawn"
safeInfoText.TextColor3 = Theme.TextDim
safeInfoText.TextSize = FS(9)
safeInfoText.Font = Enum.Font.Gotham
safeInfoText.TextWrapped = true
safeInfoText.TextXAlignment = Enum.TextXAlignment.Left
safeInfoText.Parent = safeInfoCard

-- Status card with live info
local safeStatusCard = MakeCard(GP, IsMobile and 80 or 75)

local SafeTPStatusLabel = Instance.new("TextLabel")
SafeTPStatusLabel.Size = UDim2.new(1, -12, 0, 18)
SafeTPStatusLabel.Position = UDim2.new(0, 6, 0, 4)
SafeTPStatusLabel.BackgroundTransparency = 1
SafeTPStatusLabel.Text = "🛡️ Ready to Safe TP"
SafeTPStatusLabel.TextColor3 = Theme.TextDim
SafeTPStatusLabel.TextSize = FS(11)
SafeTPStatusLabel.Font = Enum.Font.GothamBold
SafeTPStatusLabel.TextXAlignment = Enum.TextXAlignment.Left
SafeTPStatusLabel.Parent = safeStatusCard

local SafeTPPhaseLabel = Instance.new("TextLabel")
SafeTPPhaseLabel.Size = UDim2.new(1, -12, 0, 14)
SafeTPPhaseLabel.Position = UDim2.new(0, 6, 0, 24)
SafeTPPhaseLabel.BackgroundTransparency = 1
SafeTPPhaseLabel.Text = ""
SafeTPPhaseLabel.TextColor3 = Theme.TextMuted
SafeTPPhaseLabel.TextSize = FS(9)
SafeTPPhaseLabel.Font = Enum.Font.Gotham
SafeTPPhaseLabel.TextXAlignment = Enum.TextXAlignment.Left
SafeTPPhaseLabel.Parent = safeStatusCard

-- Progress bar background
local progressBg = Instance.new("Frame")
progressBg.Size = UDim2.new(1, -16, 0, IsMobile and 14 or 10)
progressBg.Position = UDim2.new(0, 8, 0, IsMobile and 46 or 44)
progressBg.BackgroundColor3 = Theme.BgLighter
progressBg.BorderSizePixel = 0
progressBg.Parent = safeStatusCard
Corner(progressBg, IsMobile and 7 or 5)

local SafeTPProgress = Instance.new("Frame")
SafeTPProgress.Size = UDim2.new(0, 0, 1, 0)
SafeTPProgress.BackgroundColor3 = Theme.Cyan
SafeTPProgress.BorderSizePixel = 0
SafeTPProgress.Parent = progressBg
Corner(SafeTPProgress, IsMobile and 7 or 5)

-- Animate progress bar color
task.spawn(function()
    local hue = 0.5
    while SafeTPProgress and SafeTPProgress.Parent do
        if IsSafeTPing then
            hue = (hue + 0.01) % 1
            SafeTPProgress.BackgroundColor3 = Color3.fromHSV(hue, 0.8, 1)
        end
        task.wait(0.03)
    end
end)

local SafeTPRouteLabel = Instance.new("TextLabel")
SafeTPRouteLabel.Size = UDim2.new(1, -12, 0, 12)
SafeTPRouteLabel.Position = UDim2.new(0, 6, 0, IsMobile and 62 or 58)
SafeTPRouteLabel.BackgroundTransparency = 1
SafeTPRouteLabel.Text = "📍 Spawn: (73, -3, 0) | Depth: Y=-42 (Safe)"
SafeTPRouteLabel.TextColor3 = Theme.TextMuted
SafeTPRouteLabel.TextSize = FS(8)
SafeTPRouteLabel.Font = Enum.Font.Gotham
SafeTPRouteLabel.TextXAlignment = Enum.TextXAlignment.Left
SafeTPRouteLabel.Parent = safeStatusCard

-- Speed & Depth sliders
MakeSlider(GP, "⚡ Horizontal Speed", 50, 500, 200, " st/s", function(val)
    SafeTPSpeed = val
end)

MakeSlider(GP, "⬆️ Vertical Speed", 50, 400, 150, " st/s", function(val)
    SafeTPVertSpeed = val
end)

MakeSlider(GP, "⬇️ Underground Depth", -200, -20, -42, " Y", function(val)
    UndergroundY = val
    if SafeTPRouteLabel and SafeTPRouteLabel.Parent then
        SafeTPRouteLabel.Text = "📍 Spawn: (73, -3, 0) | Depth: Y=" .. val
    end
end)

-- Main Safe TP button
local safeTPBtnCard = Instance.new("Frame")
safeTPBtnCard.Size = UDim2.new(1, 0, 0, IsMobile and 46 or 42)
safeTPBtnCard.BackgroundTransparency = 1
safeTPBtnCard.Parent = GP

local safeGoBtn = Instance.new("TextButton")
safeGoBtn.Size = UDim2.new(0.65, -4, 1, 0)
safeGoBtn.Position = UDim2.new(0, 0, 0, 0)
safeGoBtn.BackgroundColor3 = Theme.Cyan
safeGoBtn.Text = "🛡️ SAFE TP TO SPAWN"
safeGoBtn.TextColor3 = Theme.Bg
safeGoBtn.TextSize = FS(IsMobile and 12 or 11)
safeGoBtn.Font = Enum.Font.GothamBlack
safeGoBtn.BorderSizePixel = 0
safeGoBtn.AutoButtonColor = false
safeGoBtn.Parent = safeTPBtnCard
Corner(safeGoBtn, 8)
Ripple(safeGoBtn)

-- Cancel button
local safeCancelBtn = Instance.new("TextButton")
safeCancelBtn.Size = UDim2.new(0.35, -4, 1, 0)
safeCancelBtn.Position = UDim2.new(0.65, 4, 0, 0)
safeCancelBtn.BackgroundColor3 = Theme.Red
safeCancelBtn.Text = "⛔ CANCEL"
safeCancelBtn.TextColor3 = Theme.Text
safeCancelBtn.TextSize = FS(IsMobile and 11 or 10)
safeCancelBtn.Font = Enum.Font.GothamBlack
safeCancelBtn.BorderSizePixel = 0
safeCancelBtn.AutoButtonColor = false
safeCancelBtn.Parent = safeTPBtnCard
Corner(safeCancelBtn, 8)
Ripple(safeCancelBtn)

safeGoBtn.MouseButton1Click:Connect(function()
    if IsSafeTPing then
        safeGoBtn.Text = "⚠️ ALREADY RUNNING!"
        task.delay(1, function() safeGoBtn.Text = "🛡️ SAFE TP TO SPAWN" end)
        return
    end

    -- Reset progress bar
    SafeTPProgress.Size = UDim2.new(0, 0, 1, 0)
    SafeTPProgress.BackgroundColor3 = Theme.Cyan

    safeGoBtn.Text = "🛡️ TELEPORTING..."
    safeGoBtn.BackgroundColor3 = Theme.Yellow

    task.spawn(function()
        SafeTPToSpawn(SafeTPStatusLabel, SafeTPPhaseLabel, SafeTPProgress)

        safeGoBtn.Text = "🛡️ SAFE TP TO SPAWN"
        safeGoBtn.BackgroundColor3 = Theme.Cyan
    end)
end)

safeCancelBtn.MouseButton1Click:Connect(function()
    if IsSafeTPing then
        SafeTPCancel = true
        safeCancelBtn.Text = "⛔ STOPPING..."
        task.delay(1.5, function() safeCancelBtn.Text = "⛔ CANCEL" end)
    end
end)

-- Instant (unsafe) TP to spawn for comparison
MakeButton(GP, "⚡ INSTANT TP TO SPAWN (UNSAFE)", Theme.Orange, function(btn)
    local hrp = GetHRP()
    if hrp then
        hrp.CFrame = SpawnCFrame
        pcall(function()
            hrp.Velocity = Vector3.zero
            hrp.AssemblyLinearVelocity = Vector3.zero
        end)
        btn.Text = "✅ TELEPORTED!"
        LastTP = SpawnCFrame
        LastTPLabel.Text = "🎯 Spawn"
    else
        btn.Text = "❌ FAILED"
    end
    task.delay(1.5, function() btn.Text = "⚡ INSTANT TP TO SPAWN (UNSAFE)" end)
end)

Instance.new("Frame", GP).Size = UDim2.new(1, 0, 0, 8)

-- ════════════════════════════════════════
-- // OTHER PAGES
-- ════════════════════════════════════════

-- COINS PAGE
local CP = Pages["Coins"]
MakeRainbowLabel(CP, "🪙 COIN SCANNER", 12)

MakeToggle(CP, "🪙 Auto Scan & Collect", false, function(state)
    AutoCoinScan = state
    if state then StartCoinScanner() else StopCoinScanner() end
end)

MakeSlider(CP, "📏 Scan Radius", 10, 150, 50, " studs", function(val) ScanRadius = val end)
