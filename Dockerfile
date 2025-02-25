# IMAGEN 1
# =======================================================================================================
# ====================================  STAGE 1: BASE STAGE  ============================================
# FROM    mcr.microsoft.com/dotnet/aspnet:9.0.0-bookworm-slim AS base
FROM    mcr.microsoft.com/dotnet/aspnet:8.0.11-bookworm-slim AS base
LABEL   author="maxexee"
WORKDIR /app
EXPOSE  80

# =======================================================================================================
# =======================================================================================================
# ====================================  STAGE 2: BUILD STAGE  ===========================================
# FROM    mcr.microsoft.com/dotnet/sdk:9.0.101-bookworm-slim AS  build
FROM    mcr.microsoft.com/dotnet/sdk:8.0.404-bookworm-slim AS  build
WORKDIR /src
COPY    ["RuGa.csproj", "./"]
RUN     dotnet  restore "RuGa.csproj"
COPY    .   .
RUN     dotnet build   "RuGa.csproj" -c Release -o /app/build

# =======================================================================================================
# =======================================================================================================
#====================================   STAGE 3: PUBLISH STAGE  =========================================
FROM    build   AS  publish
RUN     dotnet  publish "RuGa.csproj"   -c Release  -o  /app/publish --self-contained true

# =======================================================================================================
# =======================================================================================================
#====================================   STAGE 4: RUN STAGE  =============================================
FROM    base    AS  final
WORKDIR /app
COPY    --from=publish  /app/publish    .
ENTRYPOINT [ "dotnet",  "RuGa.dll" ]