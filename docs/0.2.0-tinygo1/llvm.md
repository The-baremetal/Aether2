
Index ¶

    Constants
    func AttributeKindID(name string) (id uint)
    func DefaultTargetTriple() (triple string)
    func InitializeAllAsmParsers()
    func InitializeAllAsmPrinters()
    func InitializeAllTargetInfos()
    func InitializeAllTargetMCs()
    func InitializeAllTargets()
    func InitializeNativeAsmPrinter() error
    func InitializeNativeTarget() error
    func LinkInInterpreter()
    func LinkInMCJIT()
    func LinkModules(Dest, Src Module) error
    func LoadLibraryPermanently(lib string) error
    func MDKindID(name string) (id int)
    func ParseCommandLineOptions(args []string, overview string)
    func VerifyFunction(f Value, a VerifierFailureAction) error
    func VerifyModule(m Module, a VerifierFailureAction) error
    func ViewFunctionCFG(f Value)
    func ViewFunctionCFGOnly(f Value)
    func WriteBitcodeToFile(m Module, file *os.File) error
    type AtomicOrdering
    type AtomicRMWBinOp
    type Attribute
        func (a Attribute) GetEnumKind() (id int)
        func (a Attribute) GetEnumValue() (val uint64)
        func (a Attribute) GetStringKind() string
        func (a Attribute) GetStringValue() string
        func (a Attribute) GetTypeValue() (t Type)
        func (a Attribute) IsEnum() bool
        func (c Attribute) IsNil() bool
        func (a Attribute) IsString() bool
    type BasicBlock
        func AddBasicBlock(f Value, name string) (bb BasicBlock)
        func InsertBasicBlock(ref BasicBlock, name string) (bb BasicBlock)
        func NextBasicBlock(bb BasicBlock) (rbb BasicBlock)
        func PrevBasicBlock(bb BasicBlock) (rbb BasicBlock)
        func (bb BasicBlock) AsValue() (v Value)
        func (bb BasicBlock) EraseFromParent()
        func (bb BasicBlock) FirstInstruction() (v Value)
        func (c BasicBlock) IsNil() bool
        func (bb BasicBlock) LastInstruction() (v Value)
        func (bb BasicBlock) MoveAfter(pos BasicBlock)
        func (bb BasicBlock) MoveBefore(pos BasicBlock)
        func (bb BasicBlock) Parent() (v Value)
    type Builder
        func (b Builder) ClearInsertionPoint()
        func (b Builder) CreateAShr(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateAdd(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateAggregateRet(vs []Value) (rv Value)
        func (b Builder) CreateAlloca(t Type, name string) (v Value)
        func (b Builder) CreateAnd(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateArrayAlloca(t Type, val Value, name string) (v Value)
        func (b Builder) CreateArrayMalloc(t Type, val Value, name string) (v Value)
        func (b Builder) CreateAtomicCmpXchg(ptr, cmp, newVal Value, successOrdering, failureOrdering AtomicOrdering, ...) (v Value)
        func (b Builder) CreateAtomicRMW(op AtomicRMWBinOp, ptr, val Value, ordering AtomicOrdering, singleThread bool) (v Value)
        func (b Builder) CreateBinOp(op Opcode, lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateBitCast(val Value, t Type, name string) (v Value)
        func (b Builder) CreateBr(bb BasicBlock) (rv Value)
        func (b Builder) CreateCall(t Type, fn Value, args []Value, name string) (v Value)
        func (b Builder) CreateCast(val Value, op Opcode, t Type, name string) (v Value)
        func (b Builder) CreateCatchPad(parentPad Value, args []Value, name string) (v Value)
        func (b Builder) CreateCatchRet(catchpad Value, bb BasicBlock) (v Value)
        func (b Builder) CreateCatchSwitch(parentPad Value, unwindBB BasicBlock, numHandlers int, name string) (v Value)
        func (b Builder) CreateCleanupPad(parentPad Value, args []Value, name string) (v Value)
        func (b Builder) CreateCleanupRet(catchpad Value, bb BasicBlock) (v Value)
        func (b Builder) CreateCondBr(ifv Value, thenb, elseb BasicBlock) (rv Value)
        func (b Builder) CreateExactSDiv(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateExtractElement(vec, i Value, name string) (v Value)
        func (b Builder) CreateExtractValue(agg Value, i int, name string) (v Value)
        func (b Builder) CreateFAdd(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateFCmp(pred FloatPredicate, lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateFDiv(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateFMul(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateFNeg(v Value, name string) (rv Value)
        func (b Builder) CreateFPCast(val Value, t Type, name string) (v Value)
        func (b Builder) CreateFPExt(val Value, t Type, name string) (v Value)
        func (b Builder) CreateFPToSI(val Value, t Type, name string) (v Value)
        func (b Builder) CreateFPToUI(val Value, t Type, name string) (v Value)
        func (b Builder) CreateFPTrunc(val Value, t Type, name string) (v Value)
        func (b Builder) CreateFRem(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateFSub(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateFree(p Value) (v Value)
        func (b Builder) CreateGEP(t Type, p Value, indices []Value, name string) (v Value)
        func (b Builder) CreateGlobalString(str, name string) (v Value)
        func (b Builder) CreateGlobalStringPtr(str, name string) (v Value)
        func (b Builder) CreateICmp(pred IntPredicate, lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateInBoundsGEP(t Type, p Value, indices []Value, name string) (v Value)
        func (b Builder) CreateIndirectBr(addr Value, numDests int) (rv Value)
        func (b Builder) CreateInsertElement(vec, elt, i Value, name string) (v Value)
        func (b Builder) CreateInsertValue(agg, elt Value, i int, name string) (v Value)
        func (b Builder) CreateIntCast(val Value, t Type, name string) (v Value)
        func (b Builder) CreateIntToPtr(val Value, t Type, name string) (v Value)
        func (b Builder) CreateInvoke(t Type, fn Value, args []Value, then, catch BasicBlock, name string) (rv Value)
        func (b Builder) CreateIsNotNull(val Value, name string) (v Value)
        func (b Builder) CreateIsNull(val Value, name string) (v Value)
        func (b Builder) CreateLShr(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateLandingPad(t Type, nclauses int, name string) (l Value)
        func (b Builder) CreateLoad(t Type, p Value, name string) (v Value)
        func (b Builder) CreateMalloc(t Type, name string) (v Value)
        func (b Builder) CreateMul(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateNSWAdd(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateNSWMul(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateNSWNeg(v Value, name string) (rv Value)
        func (b Builder) CreateNSWSub(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateNUWAdd(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateNUWMul(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateNUWSub(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateNeg(v Value, name string) (rv Value)
        func (b Builder) CreateNot(v Value, name string) (rv Value)
        func (b Builder) CreateOr(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreatePHI(t Type, name string) (v Value)
        func (b Builder) CreatePointerCast(val Value, t Type, name string) (v Value)
        func (b Builder) CreatePtrDiff(t Type, lhs, rhs Value, name string) (v Value)
        func (b Builder) CreatePtrToInt(val Value, t Type, name string) (v Value)
        func (b Builder) CreateResume(ex Value) (v Value)
        func (b Builder) CreateRet(v Value) (rv Value)
        func (b Builder) CreateRetVoid() (rv Value)
        func (b Builder) CreateSDiv(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateSExt(val Value, t Type, name string) (v Value)
        func (b Builder) CreateSExtOrBitCast(val Value, t Type, name string) (v Value)
        func (b Builder) CreateSIToFP(val Value, t Type, name string) (v Value)
        func (b Builder) CreateSRem(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateSelect(ifv, thenv, elsev Value, name string) (v Value)
        func (b Builder) CreateShl(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateShuffleVector(v1, v2, mask Value, name string) (v Value)
        func (b Builder) CreateStore(val Value, p Value) (v Value)
        func (b Builder) CreateStructGEP(t Type, p Value, i int, name string) (v Value)
        func (b Builder) CreateSub(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateSwitch(v Value, elseb BasicBlock, numCases int) (rv Value)
        func (b Builder) CreateTrunc(val Value, t Type, name string) (v Value)
        func (b Builder) CreateTruncOrBitCast(val Value, t Type, name string) (v Value)
        func (b Builder) CreateUDiv(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateUIToFP(val Value, t Type, name string) (v Value)
        func (b Builder) CreateURem(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateUnreachable() (rv Value)
        func (b Builder) CreateVAArg(list Value, t Type, name string) (v Value)
        func (b Builder) CreateXor(lhs, rhs Value, name string) (v Value)
        func (b Builder) CreateZExt(val Value, t Type, name string) (v Value)
        func (b Builder) CreateZExtOrBitCast(val Value, t Type, name string) (v Value)
        func (b Builder) Dispose()
        func (b Builder) GetCurrentDebugLocation() (loc DebugLoc)
        func (b Builder) GetInsertBlock() (bb BasicBlock)
        func (b Builder) Insert(instr Value)
        func (b Builder) InsertDeclare(module Module, storage Value, md Value) Value
        func (b Builder) InsertWithName(instr Value, name string)
        func (c Builder) IsNil() bool
        func (b Builder) SetCurrentDebugLocation(line, col uint, scope, inlinedAt Metadata)
        func (b Builder) SetInsertPoint(block BasicBlock, instr Value)
        func (b Builder) SetInsertPointAtEnd(block BasicBlock)
        func (b Builder) SetInsertPointBefore(instr Value)
        func (b Builder) SetInstDebugLocation(v Value)
    type ByteOrdering
    type CallConv
    type CodeGenFileType
    type CodeGenOptLevel
    type CodeModel
    type Comdat
        func (c Comdat) SelectionKind() ComdatSelectionKind
        func (c Comdat) SetSelectionKind(k ComdatSelectionKind)
    type ComdatSelectionKind
    type Context
        func GlobalContext() Context
        func NewContext() Context
        func (c Context) AddBasicBlock(f Value, name string) (bb BasicBlock)
        func (c Context) ConstString(str string, addnull bool) (v Value)
        func (c Context) ConstStruct(constVals []Value, packed bool) (v Value)
        func (c Context) CreateEnumAttribute(kind uint, val uint64) (a Attribute)
        func (c Context) CreateStringAttribute(kind string, val string) (a Attribute)
        func (c Context) CreateTypeAttribute(kind uint, t Type) (a Attribute)
        func (c Context) Dispose()
        func (c Context) DoubleType() (t Type)
        func (c Context) FP128Type() (t Type)
        func (c Context) FloatType() (t Type)
        func (c Context) InsertBasicBlock(ref BasicBlock, name string) (bb BasicBlock)
        func (c Context) Int16Type() (t Type)
        func (c Context) Int1Type() (t Type)
        func (c Context) Int32Type() (t Type)
        func (c Context) Int64Type() (t Type)
        func (c Context) Int8Type() (t Type)
        func (c Context) IntType(numbits int) (t Type)
        func (c Context) IsNil() bool
        func (c Context) LabelType() (t Type)
        func (c Context) MDKindID(name string) (id int)
        func (c Context) MDNode(mds []Metadata) (md Metadata)
        func (c Context) MDString(str string) (md Metadata)
        func (c Context) NewBuilder() (b Builder)
        func (c Context) NewModule(name string) (m Module)
        func (c Context) PPCFP128Type() (t Type)
        func (c Context) ParseBitcodeFile(name string) (Module, error)
        func (c *Context) ParseIR(buf MemoryBuffer) (Module, error)
        func (c Context) StructCreateNamed(name string) (t Type)
        func (c Context) StructType(elementTypes []Type, packed bool) (t Type)
        func (c Context) TemporaryMDNode(mds []Metadata) (md Metadata)
        func (c Context) TokenType() (t Type)
        func (c Context) VoidType() (t Type)
        func (c Context) X86FP80Type() (t Type)
    type DIArrayType
    type DIAutoVariable
    type DIBasicType
    type DIBuilder
        func NewDIBuilder(m Module) *DIBuilder
        func (d *DIBuilder) CreateArrayType(t DIArrayType) Metadata
        func (d *DIBuilder) CreateAutoVariable(scope Metadata, v DIAutoVariable) Metadata
        func (d *DIBuilder) CreateBasicType(t DIBasicType) Metadata
        func (d *DIBuilder) CreateCompileUnit(cu DICompileUnit) Metadata
        func (d *DIBuilder) CreateExpression(addr []uint64) Metadata
        func (d *DIBuilder) CreateFile(filename, dir string) Metadata
        func (d *DIBuilder) CreateFunction(diScope Metadata, f DIFunction) Metadata
        func (d *DIBuilder) CreateGlobalVariableExpression(diScope Metadata, g DIGlobalVariableExpression) Metadata
        func (d *DIBuilder) CreateLexicalBlock(diScope Metadata, b DILexicalBlock) Metadata
        func (d *DIBuilder) CreateLexicalBlockFile(diScope Metadata, diFile Metadata, discriminator int) Metadata
        func (d *DIBuilder) CreateMemberType(scope Metadata, t DIMemberType) Metadata
        func (d *DIBuilder) CreateParameterVariable(scope Metadata, v DIParameterVariable) Metadata
        func (d *DIBuilder) CreatePointerType(t DIPointerType) Metadata
        func (d *DIBuilder) CreateReplaceableCompositeType(scope Metadata, t DIReplaceableCompositeType) Metadata
        func (d *DIBuilder) CreateStructType(scope Metadata, t DIStructType) Metadata
        func (d *DIBuilder) CreateSubroutineType(t DISubroutineType) Metadata
        func (d *DIBuilder) CreateTypedef(t DITypedef) Metadata
        func (d *DIBuilder) Destroy()
        func (d *DIBuilder) Finalize()
        func (d *DIBuilder) InsertValueAtEnd(v Value, diVarInfo, expr Metadata, l DebugLoc, bb BasicBlock)
    type DICompileUnit
    type DIFunction
    type DIGlobalVariableExpression
    type DILexicalBlock
    type DIMemberType
    type DIParameterVariable
    type DIPointerType
    type DIReplaceableCompositeType
    type DIStructType
    type DISubrange
    type DISubroutineType
    type DITypedef
    type DebugLoc
    type DwarfLang
    type DwarfTag
    type DwarfTypeEncoding
    type ExecutionEngine
        func NewExecutionEngine(m Module) (ee ExecutionEngine, err error)
        func NewInterpreter(m Module) (ee ExecutionEngine, err error)
        func NewMCJITCompiler(m Module, options MCJITCompilerOptions) (ee ExecutionEngine, err error)
        func (ee ExecutionEngine) AddGlobalMapping(global Value, addr unsafe.Pointer)
        func (ee ExecutionEngine) AddModule(m Module)
        func (ee ExecutionEngine) Dispose()
        func (ee ExecutionEngine) FindFunction(name string) (f Value)
        func (ee ExecutionEngine) FreeMachineCodeForFunction(f Value)
        func (ee ExecutionEngine) PointerToGlobal(global Value) unsafe.Pointer
        func (ee ExecutionEngine) RecompileAndRelinkFunction(f Value) unsafe.Pointer
        func (ee ExecutionEngine) RemoveModule(m Module)
        func (ee ExecutionEngine) RunFunction(f Value, args []GenericValue) (g GenericValue)
        func (ee ExecutionEngine) RunStaticConstructors()
        func (ee ExecutionEngine) RunStaticDestructors()
        func (ee ExecutionEngine) TargetData() (td TargetData)
    type FloatPredicate
    type GenericValue
        func NewGenericValueFromFloat(t Type, n float64) (g GenericValue)
        func NewGenericValueFromInt(t Type, n uint64, signed bool) (g GenericValue)
        func NewGenericValueFromPointer(p unsafe.Pointer) (g GenericValue)
        func (g GenericValue) Dispose()
        func (g GenericValue) Float(t Type) float64
        func (g GenericValue) Int(signed bool) uint64
        func (g GenericValue) IntWidth() int
        func (g GenericValue) Pointer() unsafe.Pointer
    type InlineAsmDialect
    type IntPredicate
    type LandingPadClause
    type Linkage
    type MCJITCompilerOptions
        func NewMCJITCompilerOptions() MCJITCompilerOptions
        func (options *MCJITCompilerOptions) SetMCJITCodeModel(CodeModel CodeModel)
        func (options *MCJITCompilerOptions) SetMCJITEnableFastISel(fastisel bool)
        func (options *MCJITCompilerOptions) SetMCJITNoFramePointerElim(nfp bool)
        func (options *MCJITCompilerOptions) SetMCJITOptimizationLevel(level uint)
    type MemoryBuffer
        func NewMemoryBufferFromFile(path string) (b MemoryBuffer, err error)
        func NewMemoryBufferFromStdin() (b MemoryBuffer, err error)
        func WriteBitcodeToMemoryBuffer(m Module) MemoryBuffer
        func WriteThinLTOBitcodeToMemoryBuffer(m Module) MemoryBuffer
        func (b MemoryBuffer) Bytes() []byte
        func (b MemoryBuffer) Dispose()
        func (c MemoryBuffer) IsNil() bool
    type Metadata
        func (md Metadata) FileDirectory() string
        func (md Metadata) FileFilename() string
        func (md Metadata) FileSource() string
        func (c Metadata) IsNil() bool
        func (md Metadata) Kind() MetadataKind
        func (md Metadata) LocationColumn() uint
        func (md Metadata) LocationInlinedAt() Metadata
        func (md Metadata) LocationLine() uint
        func (md Metadata) LocationScope() Metadata
        func (md Metadata) ReplaceAllUsesWith(new Metadata)
        func (md Metadata) ScopeFile() Metadata
        func (md Metadata) SubprogramLine() uint
    type MetadataKind
    type Module
        func (m Module) AddNamedMetadataOperand(name string, operand Metadata)
        func (m Module) Comdat(name string) (c Comdat)
        func (m Module) Context() (c Context)
        func (m Module) DataLayout() string
        func (m Module) Dispose()
        func (m Module) Dump()
        func (m Module) FirstFunction() (v Value)
        func (m Module) FirstGlobal() (v Value)
        func (m Module) GetTypeByName(name string) (t Type)
        func (c Module) IsNil() bool
        func (m Module) LastFunction() (v Value)
        func (m Module) LastGlobal() (v Value)
        func (m Module) NamedFunction(name string) (v Value)
        func (m Module) NamedGlobal(name string) (v Value)
        func (mod Module) RunPasses(passes string, tm TargetMachine, options PassBuilderOptions) error
        func (m Module) SetDataLayout(layout string)
        func (m Module) SetInlineAsm(asm string)
        func (m Module) SetTarget(target string)
        func (m Module) String() string
        func (m Module) Target() string
    type ModuleProvider
        func NewModuleProviderForModule(m Module) (mp ModuleProvider)
        func (mp ModuleProvider) Dispose()
        func (c ModuleProvider) IsNil() bool
    type Opcode
    type PassBuilderOptions
        func NewPassBuilderOptions() (pbo PassBuilderOptions)
        func (pbo PassBuilderOptions) Dispose()
        func (pbo PassBuilderOptions) SetCallGraphProfile(cgProfile bool)
        func (pbo PassBuilderOptions) SetDebugLogging(debugLogging bool)
        func (pbo PassBuilderOptions) SetForgetAllSCEVInLoopUnroll(forgetSCEV bool)
        func (pbo PassBuilderOptions) SetLicmMssaNoAccForPromotionCap(promotionCap uint)
        func (pbo PassBuilderOptions) SetLicmMssaOptCap(optCap uint)
        func (pbo PassBuilderOptions) SetLoopInterleaving(loopInterleaving bool)
        func (pbo PassBuilderOptions) SetLoopUnrolling(loopUnrolling bool)
        func (pbo PassBuilderOptions) SetLoopVectorization(loopVectorization bool)
        func (pbo PassBuilderOptions) SetMergeFunctions(mergeFuncs bool)
        func (pbo PassBuilderOptions) SetSLPVectorization(slpVectorization bool)
        func (pbo PassBuilderOptions) SetVerifyEach(verifyEach bool)
    type PassManager
        func NewFunctionPassManagerForModule(m Module) (pm PassManager)
        func NewPassManager() (pm PassManager)
        func (pm PassManager) Dispose()
        func (pm PassManager) FinalizeFunc() bool
        func (pm PassManager) InitializeFunc() bool
        func (c PassManager) IsNil() bool
        func (pm PassManager) Run(m Module) bool
        func (pm PassManager) RunFunc(f Value) bool
    type RelocMode
    type Target
        func FirstTarget() Target
        func GetTargetFromTriple(triple string) (t Target, err error)
        func (t Target) CreateTargetMachine(Triple string, CPU string, Features string, Level CodeGenOptLevel, ...) (tm TargetMachine)
        func (t Target) Description() string
        func (t Target) Name() string
        func (t Target) NextTarget() Target
    type TargetData
        func NewTargetData(rep string) (td TargetData)
        func (td TargetData) ABITypeAlignment(t Type) int
        func (td TargetData) ByteOrder() ByteOrdering
        func (td TargetData) CallFrameTypeAlignment(t Type) int
        func (td TargetData) Dispose()
        func (td TargetData) ElementContainingOffset(t Type, offset uint64) int
        func (td TargetData) ElementOffset(t Type, element int) uint64
        func (td TargetData) IntPtrType() (t Type)
        func (td TargetData) PointerSize() int
        func (td TargetData) PrefTypeAlignment(t Type) int
        func (td TargetData) PreferredAlignment(g Value) int
        func (td TargetData) String() (s string)
        func (td TargetData) TypeAllocSize(t Type) uint64
        func (td TargetData) TypeSizeInBits(t Type) uint64
        func (td TargetData) TypeStoreSize(t Type) uint64
    type TargetMachine
        func (tm TargetMachine) AddAnalysisPasses(pm PassManager)
        func (tm TargetMachine) CreateTargetData() TargetData
        func (tm TargetMachine) Dispose()
        func (tm TargetMachine) EmitToMemoryBuffer(m Module, ft CodeGenFileType) (MemoryBuffer, error)
        func (tm TargetMachine) Triple() string
    type Type
        func ArrayType(elementType Type, elementCount int) (t Type)
        func FunctionType(returnType Type, paramTypes []Type, isVarArg bool) (t Type)
        func PointerType(elementType Type, addressSpace int) (t Type)
        func StructType(elementTypes []Type, packed bool) (t Type)
        func VectorType(elementType Type, elementCount int) (t Type)
        func (t Type) ArrayLength() int
        func (t Type) Context() (c Context)
        func (t Type) ElementType() (rt Type)
        func (t Type) IntTypeWidth() int
        func (t Type) IsFunctionVarArg() bool
        func (c Type) IsNil() bool
        func (t Type) IsStructPacked() bool
        func (t Type) ParamTypes() []Type
        func (t Type) ParamTypesCount() int
        func (t Type) PointerAddressSpace() int
        func (t Type) ReturnType() (rt Type)
        func (t Type) String() string
        func (t Type) StructElementTypes() []Type
        func (t Type) StructElementTypesCount() int
        func (t Type) StructName() string
        func (t Type) StructSetBody(elementTypes []Type, packed bool)
        func (t Type) Subtypes() (ret []Type)
        func (t Type) TypeKind() TypeKind
        func (t Type) VectorSize() int
    type TypeKind
        func (t TypeKind) String() string
    type Use
        func (c Use) IsNil() bool
        func (u Use) NextUse() (ru Use)
        func (u Use) UsedValue() (v Value)
        func (u Use) User() (v Value)
    type Value
        func AddAlias(m Module, t Type, addressSpace int, aliasee Value, name string) (v Value)
        func AddFunction(m Module, name string, ft Type) (v Value)
        func AddGlobal(m Module, t Type, name string) (v Value)
        func AddGlobalInAddressSpace(m Module, t Type, name string, addressSpace int) (v Value)
        func AlignOf(t Type) (v Value)
        func BlockAddress(f Value, bb BasicBlock) (v Value)
        func ConstAdd(lhs, rhs Value) (v Value)
        func ConstAllOnes(t Type) (v Value)
        func ConstArray(t Type, constVals []Value) (v Value)
        func ConstBitCast(v Value, t Type) (rv Value)
        func ConstExtractElement(vec, i Value) (rv Value)
        func ConstFloat(t Type, n float64) (v Value)
        func ConstFloatFromString(t Type, str string) (v Value)
        func ConstGEP(t Type, v Value, indices []Value) (rv Value)
        func ConstInBoundsGEP(t Type, v Value, indices []Value) (rv Value)
        func ConstInsertElement(vec, elem, i Value) (rv Value)
        func ConstInt(t Type, n uint64, signExtend bool) (v Value)
        func ConstIntFromString(t Type, str string, radix int) (v Value)
        func ConstIntToPtr(v Value, t Type) (rv Value)
        func ConstMul(lhs, rhs Value) (v Value)
        func ConstNSWAdd(lhs, rhs Value) (v Value)
        func ConstNSWMul(lhs, rhs Value) (v Value)
        func ConstNSWNeg(v Value) (rv Value)
        func ConstNSWSub(lhs, rhs Value) (v Value)
        func ConstNUWAdd(lhs, rhs Value) (v Value)
        func ConstNUWMul(lhs, rhs Value) (v Value)
        func ConstNUWSub(lhs, rhs Value) (v Value)
        func ConstNamedStruct(t Type, constVals []Value) (v Value)
        func ConstNeg(v Value) (rv Value)
        func ConstNot(v Value) (rv Value)
        func ConstNull(t Type) (v Value)
        func ConstPointerCast(v Value, t Type) (rv Value)
        func ConstPointerNull(t Type) (v Value)
        func ConstPtrToInt(v Value, t Type) (rv Value)
        func ConstShuffleVector(veca, vecb, mask Value) (rv Value)
        func ConstString(str string, addnull bool) (v Value)
        func ConstStruct(constVals []Value, packed bool) (v Value)
        func ConstSub(lhs, rhs Value) (v Value)
        func ConstTrunc(v Value, t Type) (rv Value)
        func ConstTruncOrBitCast(v Value, t Type) (rv Value)
        func ConstVector(scalarConstVals []Value, packed bool) (v Value)
        func ConstXor(lhs, rhs Value) (v Value)
        func InlineAsm(t Type, asmString, constraints string, hasSideEffects, isAlignStack bool, ...) (rv Value)
        func NextFunction(v Value) (rv Value)
        func NextGlobal(v Value) (rv Value)
        func NextInstruction(v Value) (rv Value)
        func NextParam(v Value) (rv Value)
        func PrevFunction(v Value) (rv Value)
        func PrevGlobal(v Value) (rv Value)
        func PrevInstruction(v Value) (rv Value)
        func PrevParam(v Value) (rv Value)
        func SizeOf(t Type) (v Value)
        func Undef(t Type) (v Value)
        func (v Value) AddAttributeAtIndex(i int, a Attribute)
        func (v Value) AddCallSiteAttribute(i int, a Attribute)
        func (v Value) AddCase(on Value, dest BasicBlock)
        func (l Value) AddClause(v Value)
        func (v Value) AddDest(dest BasicBlock)
        func (v Value) AddFunctionAttr(a Attribute)
        func (v Value) AddHandler(bb BasicBlock)
        func (v Value) AddIncoming(vals []Value, blocks []BasicBlock)
        func (v Value) AddMetadata(kind int, md Metadata)
        func (v Value) AddTargetDependentFunctionAttr(attr, value string)
        func (v Value) Alignment() int
        func (v Value) AllocatedType() (t Type)
        func (v Value) AsBasicBlock() (bb BasicBlock)
        func (v Value) BasicBlocks() []BasicBlock
        func (v Value) BasicBlocksCount() int
        func (v Value) CalledFunctionType() (t Type)
        func (v Value) CalledValue() (rv Value)
        func (v Value) CmpXchgFailureOrdering() AtomicOrdering
        func (v Value) CmpXchgSuccessOrdering() AtomicOrdering
        func (v Value) Comdat() (c Comdat)
        func (v Value) ConstGetAsString() string
        func (v Value) ConstantAsMetadata() (md Metadata)
        func (v Value) DoubleValue() (result float64, inexact bool)
        func (v Value) Dump()
        func (v Value) EntryBasicBlock() (bb BasicBlock)
        func (v Value) EraseFromParentAsFunction()
        func (v Value) EraseFromParentAsGlobal()
        func (v Value) EraseFromParentAsInstruction()
        func (v Value) FirstBasicBlock() (bb BasicBlock)
        func (v Value) FirstParam() (rv Value)
        func (v Value) FirstUse() (u Use)
        func (v Value) FloatPredicate() FloatPredicate
        func (v Value) FunctionCallConv() CallConv
        func (v Value) GC() string
        func (v Value) GEPSourceElementType() (t Type)
        func (v Value) GetCallSiteEnumAttribute(i int, kind uint) (a Attribute)
        func (v Value) GetCallSiteStringAttribute(i int, kind string) (a Attribute)
        func (v Value) GetEnumAttributeAtIndex(i int, kind uint) (a Attribute)
        func (v Value) GetEnumFunctionAttribute(kind uint) Attribute
        func (v Value) GetHandlers() []BasicBlock
        func (v Value) GetParentCatchSwitch() (rv Value)
        func (v Value) GetStringAttributeAtIndex(i int, kind string) (a Attribute)
        func (v Value) GlobalParent() (m Module)
        func (v Value) GlobalValueType() (t Type)
        func (v Value) HasMetadata() bool
        func (v Value) IncomingBlock(i int) (bb BasicBlock)
        func (v Value) IncomingCount() int
        func (v Value) IncomingValue(i int) (rv Value)
        func (v Value) Indices() []uint32
        func (v Value) Initializer() (rv Value)
        func (v Value) InstructionCallConv() CallConv
        func (v Value) InstructionDebugLoc() (md Metadata)
        func (v Value) InstructionOpcode() Opcode
        func (v Value) InstructionParent() (bb BasicBlock)
        func (v Value) InstructionSetDebugLoc(md Metadata)
        func (v Value) IntPredicate() IntPredicate
        func (v Value) IntrinsicID() int
        func (v Value) IsAAllocaInst() (rv Value)
        func (v Value) IsAArgument() (rv Value)
        func (v Value) IsABasicBlock() (rv Value)
        func (v Value) IsABinaryOperator() (rv Value)
        func (v Value) IsABitCastInst() (rv Value)
        func (v Value) IsABranchInst() (rv Value)
        func (v Value) IsACallInst() (rv Value)
        func (v Value) IsACastInst() (rv Value)
        func (v Value) IsACmpInst() (rv Value)
        func (v Value) IsAConstant() (rv Value)
        func (v Value) IsAConstantAggregateZero() (rv Value)
        func (v Value) IsAConstantArray() (rv Value)
        func (v Value) IsAConstantExpr() (rv Value)
        func (v Value) IsAConstantFP() (rv Value)
        func (v Value) IsAConstantInt() (rv Value)
        func (v Value) IsAConstantPointerNull() (rv Value)
        func (v Value) IsAConstantStruct() (rv Value)
        func (v Value) IsAConstantVector() (rv Value)
        func (v Value) IsADbgDeclareInst() (rv Value)
        func (v Value) IsADbgInfoIntrinsic() (rv Value)
        func (v Value) IsAExtractElementInst() (rv Value)
        func (v Value) IsAExtractValueInst() (rv Value)
        func (v Value) IsAFCmpInst() (rv Value)
        func (v Value) IsAFPExtInst() (rv Value)
        func (v Value) IsAFPToSIInst() (rv Value)
        func (v Value) IsAFPToUIInst() (rv Value)
        func (v Value) IsAFPTruncInst() (rv Value)
        func (v Value) IsAFunction() (rv Value)
        func (v Value) IsAGetElementPtrInst() (rv Value)
        func (v Value) IsAGlobalAlias() (rv Value)
        func (v Value) IsAGlobalValue() (rv Value)
        func (v Value) IsAGlobalVariable() (rv Value)
        func (v Value) IsAICmpInst() (rv Value)
        func (v Value) IsAInlineAsm() (rv Value)
        func (v Value) IsAInsertElementInst() (rv Value)
        func (v Value) IsAInsertValueInst() (rv Value)
        func (v Value) IsAInstruction() (rv Value)
        func (v Value) IsAIntToPtrInst() (rv Value)
        func (v Value) IsAIntrinsicInst() (rv Value)
        func (v Value) IsAInvokeInst() (rv Value)
        func (v Value) IsALoadInst() (rv Value)
        func (v Value) IsAMemCpyInst() (rv Value)
        func (v Value) IsAMemIntrinsic() (rv Value)
        func (v Value) IsAMemMoveInst() (rv Value)
        func (v Value) IsAMemSetInst() (rv Value)
        func (v Value) IsAPHINode() (rv Value)
        func (v Value) IsAPtrToIntInst() (rv Value)
        func (v Value) IsAReturnInst() (rv Value)
        func (v Value) IsASExtInst() (rv Value)
        func (v Value) IsASIToFPInst() (rv Value)
        func (v Value) IsASelectInst() (rv Value)
        func (v Value) IsAShuffleVectorInst() (rv Value)
        func (v Value) IsAStoreInst() (rv Value)
        func (v Value) IsASwitchInst() (rv Value)
        func (v Value) IsATruncInst() (rv Value)
        func (v Value) IsAUIToFPInst() (rv Value)
        func (v Value) IsAUnaryInstruction() (rv Value)
        func (v Value) IsAUndefValue() (rv Value)
        func (v Value) IsAUnreachableInst() (rv Value)
        func (v Value) IsAUser() (rv Value)
        func (v Value) IsAVAArgInst() (rv Value)
        func (v Value) IsAZExtInst() (rv Value)
        func (v Value) IsAtomicSingleThread() bool
        func (v Value) IsBasicBlock() bool
        func (v Value) IsConstant() bool
        func (v Value) IsConstantString() bool
        func (v Value) IsDeclaration() bool
        func (v Value) IsGlobalConstant() bool
        func (c Value) IsNil() bool
        func (v Value) IsNull() bool
        func (v Value) IsTailCall() bool
        func (v Value) IsThreadLocal() bool
        func (v Value) IsUndef() bool
        func (v Value) IsVolatile() bool
        func (v Value) LastBasicBlock() (bb BasicBlock)
        func (v Value) LastParam() (rv Value)
        func (v Value) Linkage() Linkage
        func (v Value) Metadata(kind int) (rv Value)
        func (v Value) Name() string
        func (v Value) Opcode() Opcode
        func (v Value) Operand(i int) (rv Value)
        func (v Value) OperandsCount() int
        func (v Value) Ordering() AtomicOrdering
        func (v Value) Param(i int) (rv Value)
        func (v Value) ParamParent() (rv Value)
        func (v Value) Params() []Value
        func (v Value) ParamsCount() int
        func (v Value) RemoveEnumAttributeAtIndex(i int, kind uint)
        func (v Value) RemoveEnumFunctionAttribute(kind uint)
        func (v Value) RemoveFromParentAsInstruction()
        func (v Value) RemoveStringAttributeAtIndex(i int, kind string)
        func (v Value) ReplaceAllUsesWith(nv Value)
        func (v Value) SExtValue() int64
        func (v Value) Section() string
        func (v Value) SetAlignment(a int)
        func (v Value) SetAtomicSingleThread(singleThread bool)
        func (l Value) SetCleanup(cleanup bool)
        func (v Value) SetCmpXchgFailureOrdering(ordering AtomicOrdering)
        func (v Value) SetCmpXchgSuccessOrdering(ordering AtomicOrdering)
        func (v Value) SetComdat(c Comdat)
        func (v Value) SetFunctionCallConv(cc CallConv)
        func (v Value) SetGC(name string)
        func (v Value) SetGlobalConstant(gc bool)
        func (v Value) SetInitializer(cv Value)
        func (v Value) SetInstrParamAlignment(i int, align int)
        func (v Value) SetInstructionCallConv(cc CallConv)
        func (v Value) SetLinkage(l Linkage)
        func (v Value) SetMetadata(kind int, node Metadata)
        func (v Value) SetName(name string)
        func (v Value) SetOperand(i int, op Value)
        func (v Value) SetOrdering(ordering AtomicOrdering)
        func (v Value) SetParamAlignment(align int)
        func (v Value) SetParentCatchSwitch(catchSwitch Value)
        func (v Value) SetPersonality(p Value)
        func (v Value) SetSection(str string)
        func (v Value) SetSubprogram(sp Metadata)
        func (v Value) SetTailCall(is bool)
        func (v Value) SetThreadLocal(tl bool)
        func (v Value) SetUnnamedAddr(ua bool)
        func (v Value) SetVisibility(vi Visibility)
        func (v Value) SetVolatile(volatile bool)
        func (v Value) String() string
        func (v Value) Subprogram() (md Metadata)
        func (v Value) Type() (t Type)
        func (v Value) Visibility() Visibility
        func (v Value) ZExtValue() uint64
    type VerifierFailureAction
    type Visibility

Constants ¶
View Source

const (
	FlagPrivate = 1 << iota
	FlagProtected
	FlagFwdDecl
	FlagAppleBlock
	FlagReserved
	FlagVirtual
	FlagArtificial
	FlagExplicit
	FlagPrototyped
	FlagObjcClassComplete
	FlagObjectPointer
	FlagVector
	FlagStaticMember
	FlagIndirectVariable
)

View Source

const (
	MDStringMetadataKind                     = C.LLVMMDStringMetadataKind
	ConstantAsMetadataMetadataKind           = C.LLVMConstantAsMetadataMetadataKind
	LocalAsMetadataMetadataKind              = C.LLVMLocalAsMetadataMetadataKind
	DistinctMDOperandPlaceholderMetadataKind = C.LLVMDistinctMDOperandPlaceholderMetadataKind
	MDTupleMetadataKind                      = C.LLVMMDTupleMetadataKind
	DILocationMetadataKind                   = C.LLVMDILocationMetadataKind
	DIExpressionMetadataKind                 = C.LLVMDIExpressionMetadataKind
	DIGlobalVariableExpressionMetadataKind   = C.LLVMDIGlobalVariableExpressionMetadataKind
	GenericDINodeMetadataKind                = C.LLVMGenericDINodeMetadataKind
	DISubrangeMetadataKind                   = C.LLVMDISubrangeMetadataKind
	DIEnumeratorMetadataKind                 = C.LLVMDIEnumeratorMetadataKind
	DIBasicTypeMetadataKind                  = C.LLVMDIBasicTypeMetadataKind
	DIDerivedTypeMetadataKind                = C.LLVMDIDerivedTypeMetadataKind
	DICompositeTypeMetadataKind              = C.LLVMDICompositeTypeMetadataKind
	DISubroutineTypeMetadataKind             = C.LLVMDISubroutineTypeMetadataKind
	DIFileMetadataKind                       = C.LLVMDIFileMetadataKind
	DICompileUnitMetadataKind                = C.LLVMDICompileUnitMetadataKind
	DISubprogramMetadataKind                 = C.LLVMDISubprogramMetadataKind
	DILexicalBlockMetadataKind               = C.LLVMDILexicalBlockMetadataKind
	DILexicalBlockFileMetadataKind           = C.LLVMDILexicalBlockFileMetadataKind
	DINamespaceMetadataKind                  = C.LLVMDINamespaceMetadataKind
	DIModuleMetadataKind                     = C.LLVMDIModuleMetadataKind
	DITemplateTypeParameterMetadataKind      = C.LLVMDITemplateTypeParameterMetadataKind
	DITemplateValueParameterMetadataKind     = C.LLVMDITemplateValueParameterMetadataKind
	DIGlobalVariableMetadataKind             = C.LLVMDIGlobalVariableMetadataKind
	DILocalVariableMetadataKind              = C.LLVMDILocalVariableMetadataKind
	DILabelMetadataKind                      = C.LLVMDILabelMetadataKind
	DIObjCPropertyMetadataKind               = C.LLVMDIObjCPropertyMetadataKind
	DIImportedEntityMetadataKind             = C.LLVMDIImportedEntityMetadataKind
	DIMacroMetadataKind                      = C.LLVMDIMacroMetadataKind
	DIMacroFileMetadataKind                  = C.LLVMDIMacroFileMetadataKind
	DICommonBlockMetadataKind                = C.LLVMDICommonBlockMetadataKind
)

View Source

const Version = C.LLVM_VERSION_STRING

Variables ¶

This section is empty.
Functions ¶
func AttributeKindID ¶

func AttributeKindID(name string) (id uint)

func DefaultTargetTriple ¶

func DefaultTargetTriple() (triple string)

func InitializeAllAsmParsers ¶

func InitializeAllAsmParsers()

func InitializeAllAsmPrinters ¶

func InitializeAllAsmPrinters()

func InitializeAllTargetInfos ¶

func InitializeAllTargetInfos()

InitializeAllTargetInfos - The main program should call this function if it wants access to all available targets that LLVM is configured to support.
func InitializeAllTargetMCs ¶

func InitializeAllTargetMCs()

func InitializeAllTargets ¶

func InitializeAllTargets()

InitializeAllTargets - The main program should call this function if it wants to link in all available targets that LLVM is configured to support.
func InitializeNativeAsmPrinter ¶

func InitializeNativeAsmPrinter() error

func InitializeNativeTarget ¶

func InitializeNativeTarget() error

InitializeNativeTarget - The main program should call this function to initialize the native target corresponding to the host. This is useful for JIT applications to ensure that the target gets linked in correctly.
func LinkInInterpreter ¶

func LinkInInterpreter()

func LinkInMCJIT ¶

func LinkInMCJIT()

func LinkModules ¶

func LinkModules(Dest, Src Module) error

func LoadLibraryPermanently ¶

func LoadLibraryPermanently(lib string) error

Loads a dynamic library such that it may be used as an LLVM plugin. See llvm::sys::DynamicLibrary::LoadLibraryPermanently.
func MDKindID ¶

func MDKindID(name string) (id int)

func ParseCommandLineOptions ¶

func ParseCommandLineOptions(args []string, overview string)

Parse the given arguments using the LLVM command line parser. See llvm::cl::ParseCommandLineOptions.
func VerifyFunction ¶

func VerifyFunction(f Value, a VerifierFailureAction) error

Verifies that a single function is valid, taking the specified action. Useful for debugging.
func VerifyModule ¶

func VerifyModule(m Module, a VerifierFailureAction) error

Verifies that a module is valid, taking the specified action if not. Optionally returns a human-readable description of any invalid constructs.
func ViewFunctionCFG ¶

func ViewFunctionCFG(f Value)

Open up a ghostview window that displays the CFG of the current function. Useful for debugging.
func ViewFunctionCFGOnly ¶

func ViewFunctionCFGOnly(f Value)

func WriteBitcodeToFile ¶

func WriteBitcodeToFile(m Module, file *os.File) error

Types ¶
type AtomicOrdering ¶

type AtomicOrdering C.LLVMAtomicOrdering

const (
	AtomicOrderingNotAtomic              AtomicOrdering = C.LLVMAtomicOrderingNotAtomic
	AtomicOrderingUnordered              AtomicOrdering = C.LLVMAtomicOrderingUnordered
	AtomicOrderingMonotonic              AtomicOrdering = C.LLVMAtomicOrderingMonotonic
	AtomicOrderingAcquire                AtomicOrdering = C.LLVMAtomicOrderingAcquire
	AtomicOrderingRelease                AtomicOrdering = C.LLVMAtomicOrderingRelease
	AtomicOrderingAcquireRelease         AtomicOrdering = C.LLVMAtomicOrderingAcquireRelease
	AtomicOrderingSequentiallyConsistent AtomicOrdering = C.LLVMAtomicOrderingSequentiallyConsistent
)

type AtomicRMWBinOp ¶

type AtomicRMWBinOp C.LLVMAtomicRMWBinOp

const (
	AtomicRMWBinOpXchg AtomicRMWBinOp = C.LLVMAtomicRMWBinOpXchg
	AtomicRMWBinOpAdd  AtomicRMWBinOp = C.LLVMAtomicRMWBinOpAdd
	AtomicRMWBinOpSub  AtomicRMWBinOp = C.LLVMAtomicRMWBinOpSub
	AtomicRMWBinOpAnd  AtomicRMWBinOp = C.LLVMAtomicRMWBinOpAnd
	AtomicRMWBinOpNand AtomicRMWBinOp = C.LLVMAtomicRMWBinOpNand
	AtomicRMWBinOpOr   AtomicRMWBinOp = C.LLVMAtomicRMWBinOpOr
	AtomicRMWBinOpXor  AtomicRMWBinOp = C.LLVMAtomicRMWBinOpXor
	AtomicRMWBinOpMax  AtomicRMWBinOp = C.LLVMAtomicRMWBinOpMax
	AtomicRMWBinOpMin  AtomicRMWBinOp = C.LLVMAtomicRMWBinOpMin
	AtomicRMWBinOpUMax AtomicRMWBinOp = C.LLVMAtomicRMWBinOpUMax
	AtomicRMWBinOpUMin AtomicRMWBinOp = C.LLVMAtomicRMWBinOpUMin
)

type Attribute ¶

type Attribute struct {
	C C.LLVMAttributeRef
}

func (Attribute) GetEnumKind ¶

func (a Attribute) GetEnumKind() (id int)

func (Attribute) GetEnumValue ¶

func (a Attribute) GetEnumValue() (val uint64)

func (Attribute) GetStringKind ¶

func (a Attribute) GetStringKind() string

func (Attribute) GetStringValue ¶

func (a Attribute) GetStringValue() string

func (Attribute) GetTypeValue ¶

func (a Attribute) GetTypeValue() (t Type)

func (Attribute) IsEnum ¶

func (a Attribute) IsEnum() bool

func (Attribute) IsNil ¶

func (c Attribute) IsNil() bool

func (Attribute) IsString ¶

func (a Attribute) IsString() bool

type BasicBlock ¶

type BasicBlock struct {
	C C.LLVMBasicBlockRef
}

func AddBasicBlock ¶

func AddBasicBlock(f Value, name string) (bb BasicBlock)

func InsertBasicBlock ¶

func InsertBasicBlock(ref BasicBlock, name string) (bb BasicBlock)

func NextBasicBlock ¶

func NextBasicBlock(bb BasicBlock) (rbb BasicBlock)

func PrevBasicBlock ¶

func PrevBasicBlock(bb BasicBlock) (rbb BasicBlock)

func (BasicBlock) AsValue ¶

func (bb BasicBlock) AsValue() (v Value)

Operations on basic blocks
func (BasicBlock) EraseFromParent ¶

func (bb BasicBlock) EraseFromParent()

func (BasicBlock) FirstInstruction ¶

func (bb BasicBlock) FirstInstruction() (v Value)

func (BasicBlock) IsNil ¶

func (c BasicBlock) IsNil() bool

func (BasicBlock) LastInstruction ¶

func (bb BasicBlock) LastInstruction() (v Value)

func (BasicBlock) MoveAfter ¶

func (bb BasicBlock) MoveAfter(pos BasicBlock)

func (BasicBlock) MoveBefore ¶

func (bb BasicBlock) MoveBefore(pos BasicBlock)

func (BasicBlock) Parent ¶

func (bb BasicBlock) Parent() (v Value)

type Builder ¶

type Builder struct {
	C C.LLVMBuilderRef
}

func (Builder) ClearInsertionPoint ¶

func (b Builder) ClearInsertionPoint()

func (Builder) CreateAShr ¶

func (b Builder) CreateAShr(lhs, rhs Value, name string) (v Value)

func (Builder) CreateAdd ¶

func (b Builder) CreateAdd(lhs, rhs Value, name string) (v Value)

Arithmetic
func (Builder) CreateAggregateRet ¶

func (b Builder) CreateAggregateRet(vs []Value) (rv Value)

func (Builder) CreateAlloca ¶

func (b Builder) CreateAlloca(t Type, name string) (v Value)

func (Builder) CreateAnd ¶

func (b Builder) CreateAnd(lhs, rhs Value, name string) (v Value)

func (Builder) CreateArrayAlloca ¶

func (b Builder) CreateArrayAlloca(t Type, val Value, name string) (v Value)

func (Builder) CreateArrayMalloc ¶

func (b Builder) CreateArrayMalloc(t Type, val Value, name string) (v Value)

func (Builder) CreateAtomicCmpXchg ¶

func (b Builder) CreateAtomicCmpXchg(ptr, cmp, newVal Value, successOrdering, failureOrdering AtomicOrdering, singleThread bool) (v Value)

func (Builder) CreateAtomicRMW ¶

func (b Builder) CreateAtomicRMW(op AtomicRMWBinOp, ptr, val Value, ordering AtomicOrdering, singleThread bool) (v Value)

func (Builder) CreateBinOp ¶

func (b Builder) CreateBinOp(op Opcode, lhs, rhs Value, name string) (v Value)

func (Builder) CreateBitCast ¶

func (b Builder) CreateBitCast(val Value, t Type, name string) (v Value)

func (Builder) CreateBr ¶

func (b Builder) CreateBr(bb BasicBlock) (rv Value)

func (Builder) CreateCall ¶

func (b Builder) CreateCall(t Type, fn Value, args []Value, name string) (v Value)

func (Builder) CreateCast ¶

func (b Builder) CreateCast(val Value, op Opcode, t Type, name string) (v Value)

func (Builder) CreateCatchPad ¶

func (b Builder) CreateCatchPad(parentPad Value, args []Value, name string) (v Value)

func (Builder) CreateCatchRet ¶

func (b Builder) CreateCatchRet(catchpad Value, bb BasicBlock) (v Value)

func (Builder) CreateCatchSwitch ¶

func (b Builder) CreateCatchSwitch(parentPad Value, unwindBB BasicBlock, numHandlers int, name string) (v Value)

func (Builder) CreateCleanupPad ¶

func (b Builder) CreateCleanupPad(parentPad Value, args []Value, name string) (v Value)

func (Builder) CreateCleanupRet ¶

func (b Builder) CreateCleanupRet(catchpad Value, bb BasicBlock) (v Value)

func (Builder) CreateCondBr ¶

func (b Builder) CreateCondBr(ifv Value, thenb, elseb BasicBlock) (rv Value)

func (Builder) CreateExactSDiv ¶

func (b Builder) CreateExactSDiv(lhs, rhs Value, name string) (v Value)

func (Builder) CreateExtractElement ¶

func (b Builder) CreateExtractElement(vec, i Value, name string) (v Value)

func (Builder) CreateExtractValue ¶

func (b Builder) CreateExtractValue(agg Value, i int, name string) (v Value)

func (Builder) CreateFAdd ¶

func (b Builder) CreateFAdd(lhs, rhs Value, name string) (v Value)

func (Builder) CreateFCmp ¶

func (b Builder) CreateFCmp(pred FloatPredicate, lhs, rhs Value, name string) (v Value)

func (Builder) CreateFDiv ¶

func (b Builder) CreateFDiv(lhs, rhs Value, name string) (v Value)

func (Builder) CreateFMul ¶

func (b Builder) CreateFMul(lhs, rhs Value, name string) (v Value)

func (Builder) CreateFNeg ¶

func (b Builder) CreateFNeg(v Value, name string) (rv Value)

func (Builder) CreateFPCast ¶

func (b Builder) CreateFPCast(val Value, t Type, name string) (v Value)

func (Builder) CreateFPExt ¶

func (b Builder) CreateFPExt(val Value, t Type, name string) (v Value)

func (Builder) CreateFPToSI ¶

func (b Builder) CreateFPToSI(val Value, t Type, name string) (v Value)

func (Builder) CreateFPToUI ¶

func (b Builder) CreateFPToUI(val Value, t Type, name string) (v Value)

func (Builder) CreateFPTrunc ¶

func (b Builder) CreateFPTrunc(val Value, t Type, name string) (v Value)

func (Builder) CreateFRem ¶

func (b Builder) CreateFRem(lhs, rhs Value, name string) (v Value)

func (Builder) CreateFSub ¶

func (b Builder) CreateFSub(lhs, rhs Value, name string) (v Value)

func (Builder) CreateFree ¶

func (b Builder) CreateFree(p Value) (v Value)

func (Builder) CreateGEP ¶

func (b Builder) CreateGEP(t Type, p Value, indices []Value, name string) (v Value)

func (Builder) CreateGlobalString ¶

func (b Builder) CreateGlobalString(str, name string) (v Value)

func (Builder) CreateGlobalStringPtr ¶

func (b Builder) CreateGlobalStringPtr(str, name string) (v Value)

func (Builder) CreateICmp ¶

func (b Builder) CreateICmp(pred IntPredicate, lhs, rhs Value, name string) (v Value)

Comparisons
func (Builder) CreateInBoundsGEP ¶

func (b Builder) CreateInBoundsGEP(t Type, p Value, indices []Value, name string) (v Value)

func (Builder) CreateIndirectBr ¶

func (b Builder) CreateIndirectBr(addr Value, numDests int) (rv Value)

func (Builder) CreateInsertElement ¶

func (b Builder) CreateInsertElement(vec, elt, i Value, name string) (v Value)

func (Builder) CreateInsertValue ¶

func (b Builder) CreateInsertValue(agg, elt Value, i int, name string) (v Value)

func (Builder) CreateIntCast ¶

func (b Builder) CreateIntCast(val Value, t Type, name string) (v Value)

func (Builder) CreateIntToPtr ¶

func (b Builder) CreateIntToPtr(val Value, t Type, name string) (v Value)

func (Builder) CreateInvoke ¶

func (b Builder) CreateInvoke(t Type, fn Value, args []Value, then, catch BasicBlock, name string) (rv Value)

func (Builder) CreateIsNotNull ¶

func (b Builder) CreateIsNotNull(val Value, name string) (v Value)

func (Builder) CreateIsNull ¶

func (b Builder) CreateIsNull(val Value, name string) (v Value)

func (Builder) CreateLShr ¶

func (b Builder) CreateLShr(lhs, rhs Value, name string) (v Value)

func (Builder) CreateLandingPad ¶

func (b Builder) CreateLandingPad(t Type, nclauses int, name string) (l Value)

func (Builder) CreateLoad ¶

func (b Builder) CreateLoad(t Type, p Value, name string) (v Value)

func (Builder) CreateMalloc ¶

func (b Builder) CreateMalloc(t Type, name string) (v Value)

func (Builder) CreateMul ¶

func (b Builder) CreateMul(lhs, rhs Value, name string) (v Value)

func (Builder) CreateNSWAdd ¶

func (b Builder) CreateNSWAdd(lhs, rhs Value, name string) (v Value)

func (Builder) CreateNSWMul ¶

func (b Builder) CreateNSWMul(lhs, rhs Value, name string) (v Value)

func (Builder) CreateNSWNeg ¶

func (b Builder) CreateNSWNeg(v Value, name string) (rv Value)

func (Builder) CreateNSWSub ¶

func (b Builder) CreateNSWSub(lhs, rhs Value, name string) (v Value)

func (Builder) CreateNUWAdd ¶

func (b Builder) CreateNUWAdd(lhs, rhs Value, name string) (v Value)

func (Builder) CreateNUWMul ¶

func (b Builder) CreateNUWMul(lhs, rhs Value, name string) (v Value)

func (Builder) CreateNUWSub ¶

func (b Builder) CreateNUWSub(lhs, rhs Value, name string) (v Value)

func (Builder) CreateNeg ¶

func (b Builder) CreateNeg(v Value, name string) (rv Value)

func (Builder) CreateNot ¶

func (b Builder) CreateNot(v Value, name string) (rv Value)

func (Builder) CreateOr ¶

func (b Builder) CreateOr(lhs, rhs Value, name string) (v Value)

func (Builder) CreatePHI ¶

func (b Builder) CreatePHI(t Type, name string) (v Value)

Miscellaneous instructions
func (Builder) CreatePointerCast ¶

func (b Builder) CreatePointerCast(val Value, t Type, name string) (v Value)

func (Builder) CreatePtrDiff ¶

func (b Builder) CreatePtrDiff(t Type, lhs, rhs Value, name string) (v Value)

func (Builder) CreatePtrToInt ¶

func (b Builder) CreatePtrToInt(val Value, t Type, name string) (v Value)

func (Builder) CreateResume ¶

func (b Builder) CreateResume(ex Value) (v Value)

func (Builder) CreateRet ¶

func (b Builder) CreateRet(v Value) (rv Value)

func (Builder) CreateRetVoid ¶

func (b Builder) CreateRetVoid() (rv Value)

Terminators
func (Builder) CreateSDiv ¶

func (b Builder) CreateSDiv(lhs, rhs Value, name string) (v Value)

func (Builder) CreateSExt ¶

func (b Builder) CreateSExt(val Value, t Type, name string) (v Value)

func (Builder) CreateSExtOrBitCast ¶

func (b Builder) CreateSExtOrBitCast(val Value, t Type, name string) (v Value)

func (Builder) CreateSIToFP ¶

func (b Builder) CreateSIToFP(val Value, t Type, name string) (v Value)

func (Builder) CreateSRem ¶

func (b Builder) CreateSRem(lhs, rhs Value, name string) (v Value)

func (Builder) CreateSelect ¶

func (b Builder) CreateSelect(ifv, thenv, elsev Value, name string) (v Value)

func (Builder) CreateShl ¶

func (b Builder) CreateShl(lhs, rhs Value, name string) (v Value)

func (Builder) CreateShuffleVector ¶

func (b Builder) CreateShuffleVector(v1, v2, mask Value, name string) (v Value)

func (Builder) CreateStore ¶

func (b Builder) CreateStore(val Value, p Value) (v Value)

func (Builder) CreateStructGEP ¶

func (b Builder) CreateStructGEP(t Type, p Value, i int, name string) (v Value)

func (Builder) CreateSub ¶

func (b Builder) CreateSub(lhs, rhs Value, name string) (v Value)

func (Builder) CreateSwitch ¶

func (b Builder) CreateSwitch(v Value, elseb BasicBlock, numCases int) (rv Value)

func (Builder) CreateTrunc ¶

func (b Builder) CreateTrunc(val Value, t Type, name string) (v Value)

Casts
func (Builder) CreateTruncOrBitCast ¶

func (b Builder) CreateTruncOrBitCast(val Value, t Type, name string) (v Value)

func (Builder) CreateUDiv ¶

func (b Builder) CreateUDiv(lhs, rhs Value, name string) (v Value)

func (Builder) CreateUIToFP ¶

func (b Builder) CreateUIToFP(val Value, t Type, name string) (v Value)

func (Builder) CreateURem ¶

func (b Builder) CreateURem(lhs, rhs Value, name string) (v Value)

func (Builder) CreateUnreachable ¶

func (b Builder) CreateUnreachable() (rv Value)

func (Builder) CreateVAArg ¶

func (b Builder) CreateVAArg(list Value, t Type, name string) (v Value)

func (Builder) CreateXor ¶

func (b Builder) CreateXor(lhs, rhs Value, name string) (v Value)

func (Builder) CreateZExt ¶

func (b Builder) CreateZExt(val Value, t Type, name string) (v Value)

func (Builder) CreateZExtOrBitCast ¶

func (b Builder) CreateZExtOrBitCast(val Value, t Type, name string) (v Value)

func (Builder) Dispose ¶

func (b Builder) Dispose()

func (Builder) GetCurrentDebugLocation ¶

func (b Builder) GetCurrentDebugLocation() (loc DebugLoc)

Get current debug location. Please do not call this function until setting debug location with SetCurrentDebugLocation()
func (Builder) GetInsertBlock ¶

func (b Builder) GetInsertBlock() (bb BasicBlock)

func (Builder) Insert ¶

func (b Builder) Insert(instr Value)

func (Builder) InsertDeclare ¶

func (b Builder) InsertDeclare(module Module, storage Value, md Value) Value

func (Builder) InsertWithName ¶

func (b Builder) InsertWithName(instr Value, name string)

func (Builder) IsNil ¶

func (c Builder) IsNil() bool

func (Builder) SetCurrentDebugLocation ¶

func (b Builder) SetCurrentDebugLocation(line, col uint, scope, inlinedAt Metadata)

func (Builder) SetInsertPoint ¶

func (b Builder) SetInsertPoint(block BasicBlock, instr Value)

func (Builder) SetInsertPointAtEnd ¶

func (b Builder) SetInsertPointAtEnd(block BasicBlock)

func (Builder) SetInsertPointBefore ¶

func (b Builder) SetInsertPointBefore(instr Value)

func (Builder) SetInstDebugLocation ¶

func (b Builder) SetInstDebugLocation(v Value)

type ByteOrdering ¶

type ByteOrdering C.enum_LLVMByteOrdering

const (
	BigEndian    ByteOrdering = C.LLVMBigEndian
	LittleEndian ByteOrdering = C.LLVMLittleEndian
)

type CallConv ¶

type CallConv C.LLVMCallConv

const (
	CCallConv           CallConv = C.LLVMCCallConv
	FastCallConv        CallConv = C.LLVMFastCallConv
	ColdCallConv        CallConv = C.LLVMColdCallConv
	X86StdcallCallConv  CallConv = C.LLVMX86StdcallCallConv
	X86FastcallCallConv CallConv = C.LLVMX86FastcallCallConv
)

type CodeGenFileType ¶

type CodeGenFileType C.LLVMCodeGenFileType

const (
	AssemblyFile CodeGenFileType = C.LLVMAssemblyFile
	ObjectFile   CodeGenFileType = C.LLVMObjectFile
)

type CodeGenOptLevel ¶

type CodeGenOptLevel C.LLVMCodeGenOptLevel

const (
	CodeGenLevelNone       CodeGenOptLevel = C.LLVMCodeGenLevelNone
	CodeGenLevelLess       CodeGenOptLevel = C.LLVMCodeGenLevelLess
	CodeGenLevelDefault    CodeGenOptLevel = C.LLVMCodeGenLevelDefault
	CodeGenLevelAggressive CodeGenOptLevel = C.LLVMCodeGenLevelAggressive
)

type CodeModel ¶

type CodeModel C.LLVMCodeModel

const (
	CodeModelDefault    CodeModel = C.LLVMCodeModelDefault
	CodeModelJITDefault CodeModel = C.LLVMCodeModelJITDefault
	CodeModelTiny       CodeModel = C.LLVMCodeModelTiny
	CodeModelSmall      CodeModel = C.LLVMCodeModelSmall
	CodeModelKernel     CodeModel = C.LLVMCodeModelKernel
	CodeModelMedium     CodeModel = C.LLVMCodeModelMedium
	CodeModelLarge      CodeModel = C.LLVMCodeModelLarge
)

type Comdat ¶

type Comdat struct {
	C C.LLVMComdatRef
}

func (Comdat) SelectionKind ¶

func (c Comdat) SelectionKind() ComdatSelectionKind

func (Comdat) SetSelectionKind ¶

func (c Comdat) SetSelectionKind(k ComdatSelectionKind)

type ComdatSelectionKind ¶

type ComdatSelectionKind C.LLVMComdatSelectionKind

const (
	AnyComdatSelectionKind           ComdatSelectionKind = C.LLVMAnyComdatSelectionKind
	ExactMatchComdatSelectionKind    ComdatSelectionKind = C.LLVMExactMatchComdatSelectionKind
	LargestComdatSelectionKind       ComdatSelectionKind = C.LLVMLargestComdatSelectionKind
	NoDeduplicateComdatSelectionKind ComdatSelectionKind = C.LLVMNoDeduplicateComdatSelectionKind
	SameSizeComdatSelectionKind      ComdatSelectionKind = C.LLVMSameSizeComdatSelectionKind
)

type Context ¶

type Context struct {
	C C.LLVMContextRef
}

We use these weird structs here because *Ref types are pointers and Go's spec says that a pointer cannot be used as a receiver base type.
func GlobalContext ¶

func GlobalContext() Context

func NewContext ¶

func NewContext() Context

func (Context) AddBasicBlock ¶

func (c Context) AddBasicBlock(f Value, name string) (bb BasicBlock)

func (Context) ConstString ¶

func (c Context) ConstString(str string, addnull bool) (v Value)

Operations on composite constants
func (Context) ConstStruct ¶

func (c Context) ConstStruct(constVals []Value, packed bool) (v Value)

func (Context) CreateEnumAttribute ¶

func (c Context) CreateEnumAttribute(kind uint, val uint64) (a Attribute)

func (Context) CreateStringAttribute ¶

func (c Context) CreateStringAttribute(kind string, val string) (a Attribute)

func (Context) CreateTypeAttribute ¶

func (c Context) CreateTypeAttribute(kind uint, t Type) (a Attribute)

func (Context) Dispose ¶

func (c Context) Dispose()

func (Context) DoubleType ¶

func (c Context) DoubleType() (t Type)

func (Context) FP128Type ¶

func (c Context) FP128Type() (t Type)

func (Context) FloatType ¶

func (c Context) FloatType() (t Type)

Operations on real types
func (Context) InsertBasicBlock ¶

func (c Context) InsertBasicBlock(ref BasicBlock, name string) (bb BasicBlock)

func (Context) Int16Type ¶

func (c Context) Int16Type() (t Type)

func (Context) Int1Type ¶

func (c Context) Int1Type() (t Type)

Operations on integer types
func (Context) Int32Type ¶

func (c Context) Int32Type() (t Type)

func (Context) Int64Type ¶

func (c Context) Int64Type() (t Type)

func (Context) Int8Type ¶

func (c Context) Int8Type() (t Type)

func (Context) IntType ¶

func (c Context) IntType(numbits int) (t Type)

func (Context) IsNil ¶

func (c Context) IsNil() bool

func (Context) LabelType ¶

func (c Context) LabelType() (t Type)

func (Context) MDKindID ¶

func (c Context) MDKindID(name string) (id int)

func (Context) MDNode ¶

func (c Context) MDNode(mds []Metadata) (md Metadata)

func (Context) MDString ¶

func (c Context) MDString(str string) (md Metadata)

Operations on metadata
func (Context) NewBuilder ¶

func (c Context) NewBuilder() (b Builder)

func (Context) NewModule ¶

func (c Context) NewModule(name string) (m Module)

Create and destroy modules. See llvm::Module::Module.
func (Context) PPCFP128Type ¶

func (c Context) PPCFP128Type() (t Type)

func (Context) ParseBitcodeFile ¶

func (c Context) ParseBitcodeFile(name string) (Module, error)

ParseBitcodeFile parses the LLVM IR (bitcode) in the file with the specified name, and returns a new LLVM module.
func (*Context) ParseIR ¶

func (c *Context) ParseIR(buf MemoryBuffer) (Module, error)

ParseIR parses the textual IR given in the memory buffer and returns a new LLVM module in this context.
func (Context) StructCreateNamed ¶

func (c Context) StructCreateNamed(name string) (t Type)

func (Context) StructType ¶

func (c Context) StructType(elementTypes []Type, packed bool) (t Type)

Operations on struct types
func (Context) TemporaryMDNode ¶

func (c Context) TemporaryMDNode(mds []Metadata) (md Metadata)

func (Context) TokenType ¶

func (c Context) TokenType() (t Type)

func (Context) VoidType ¶

func (c Context) VoidType() (t Type)

Operations on other types
func (Context) X86FP80Type ¶

func (c Context) X86FP80Type() (t Type)

type DIArrayType ¶

type DIArrayType struct {
	SizeInBits  uint64
	AlignInBits uint32
	ElementType Metadata
	Subscripts  []DISubrange
}

DIArrayType holds the values for creating array type debug metadata.
type DIAutoVariable ¶

type DIAutoVariable struct {
	Name           string
	File           Metadata
	Line           int
	Type           Metadata
	AlwaysPreserve bool
	Flags          int
	AlignInBits    uint32
}

DIAutoVariable holds the values for creating auto variable debug metadata.
type DIBasicType ¶

type DIBasicType struct {
	Name       string
	SizeInBits uint64
	Encoding   DwarfTypeEncoding
}

DIBasicType holds the values for creating basic type debug metadata.
type DIBuilder ¶

type DIBuilder struct {
	// contains filtered or unexported fields
}

DIBuilder is a wrapper for the LLVM DIBuilder class.
func NewDIBuilder ¶

func NewDIBuilder(m Module) *DIBuilder

NewDIBuilder creates a new DIBuilder, associated with the given module.
func (*DIBuilder) CreateArrayType ¶

func (d *DIBuilder) CreateArrayType(t DIArrayType) Metadata

CreateArrayType creates struct type debug metadata.
func (*DIBuilder) CreateAutoVariable ¶

func (d *DIBuilder) CreateAutoVariable(scope Metadata, v DIAutoVariable) Metadata

CreateAutoVariable creates local variable debug metadata.
func (*DIBuilder) CreateBasicType ¶

func (d *DIBuilder) CreateBasicType(t DIBasicType) Metadata

CreateBasicType creates basic type debug metadata.
func (*DIBuilder) CreateCompileUnit ¶

func (d *DIBuilder) CreateCompileUnit(cu DICompileUnit) Metadata

CreateCompileUnit creates compile unit debug metadata.
func (*DIBuilder) CreateExpression ¶

func (d *DIBuilder) CreateExpression(addr []uint64) Metadata

CreateExpression creates a new descriptor for the specified variable which has a complex address expression for its address.
func (*DIBuilder) CreateFile ¶

func (d *DIBuilder) CreateFile(filename, dir string) Metadata

CreateFile creates file debug metadata.
func (*DIBuilder) CreateFunction ¶

func (d *DIBuilder) CreateFunction(diScope Metadata, f DIFunction) Metadata

CreateFunction creates function debug metadata.
func (*DIBuilder) CreateGlobalVariableExpression ¶

func (d *DIBuilder) CreateGlobalVariableExpression(diScope Metadata, g DIGlobalVariableExpression) Metadata

CreateGlobalVariableExpression creates a new descriptor for the specified global variable.
func (*DIBuilder) CreateLexicalBlock ¶

func (d *DIBuilder) CreateLexicalBlock(diScope Metadata, b DILexicalBlock) Metadata

CreateLexicalBlock creates lexical block debug metadata.
func (*DIBuilder) CreateLexicalBlockFile ¶

func (d *DIBuilder) CreateLexicalBlockFile(diScope Metadata, diFile Metadata, discriminator int) Metadata

func (*DIBuilder) CreateMemberType ¶

func (d *DIBuilder) CreateMemberType(scope Metadata, t DIMemberType) Metadata

CreateMemberType creates struct type debug metadata.
func (*DIBuilder) CreateParameterVariable ¶

func (d *DIBuilder) CreateParameterVariable(scope Metadata, v DIParameterVariable) Metadata

CreateParameterVariable creates parameter variable debug metadata.
func (*DIBuilder) CreatePointerType ¶

func (d *DIBuilder) CreatePointerType(t DIPointerType) Metadata

CreatePointerType creates a type that represents a pointer to another type.
func (*DIBuilder) CreateReplaceableCompositeType ¶

func (d *DIBuilder) CreateReplaceableCompositeType(scope Metadata, t DIReplaceableCompositeType) Metadata

CreateReplaceableCompositeType creates replaceable composite type debug metadata.
func (*DIBuilder) CreateStructType ¶

func (d *DIBuilder) CreateStructType(scope Metadata, t DIStructType) Metadata

CreateStructType creates struct type debug metadata.
func (*DIBuilder) CreateSubroutineType ¶

func (d *DIBuilder) CreateSubroutineType(t DISubroutineType) Metadata

CreateSubroutineType creates subroutine type debug metadata.
func (*DIBuilder) CreateTypedef ¶

func (d *DIBuilder) CreateTypedef(t DITypedef) Metadata

CreateTypedef creates typedef type debug metadata.
func (*DIBuilder) Destroy ¶

func (d *DIBuilder) Destroy()

Destroy destroys the DIBuilder.
func (*DIBuilder) Finalize ¶

func (d *DIBuilder) Finalize()

FInalize finalizes the debug information generated by the DIBuilder.
func (*DIBuilder) InsertValueAtEnd ¶

func (d *DIBuilder) InsertValueAtEnd(v Value, diVarInfo, expr Metadata, l DebugLoc, bb BasicBlock)

InsertValueAtEnd inserts a call to llvm.dbg.value at the end of the specified basic block for the given value and associated debug metadata.
type DICompileUnit ¶

type DICompileUnit struct {
	Language       DwarfLang
	File           string
	Dir            string
	Producer       string
	Optimized      bool
	Flags          string
	RuntimeVersion int
	SysRoot        string
	SDK            string
}

DICompileUnit holds the values for creating compile unit debug metadata.
type DIFunction ¶

type DIFunction struct {
	Name         string
	LinkageName  string
	File         Metadata
	Line         int
	Type         Metadata
	LocalToUnit  bool
	IsDefinition bool
	ScopeLine    int
	Flags        int
	Optimized    bool
}

DIFunction holds the values for creating function debug metadata.
type DIGlobalVariableExpression ¶

type DIGlobalVariableExpression struct {
	Name        string   // Name of the variable.
	LinkageName string   // Mangled name of the variable
	File        Metadata // File where this variable is defined.
	Line        int      // Line number.
	Type        Metadata // Variable Type.
	LocalToUnit bool     // Flag indicating whether this variable is externally visible or not.
	Expr        Metadata // The location of the global relative to the attached GlobalVariable.
	Decl        Metadata // Reference to the corresponding declaration.
	AlignInBits uint32   // Variable alignment(or 0 if no alignment attr was specified).
}

DIGlobalVariableExpression holds the values for creating global variable debug metadata.
type DILexicalBlock ¶

type DILexicalBlock struct {
	File   Metadata
	Line   int
	Column int
}

DILexicalBlock holds the values for creating lexical block debug metadata.
type DIMemberType ¶

type DIMemberType struct {
	Name         string
	File         Metadata
	Line         int
	SizeInBits   uint64
	AlignInBits  uint32
	OffsetInBits uint64
	Flags        int
	Type         Metadata
}

DIMemberType holds the values for creating member type debug metadata.
type DIParameterVariable ¶

type DIParameterVariable struct {
	Name           string
	File           Metadata
	Line           int
	Type           Metadata
	AlwaysPreserve bool
	Flags          int

	// ArgNo is the 1-based index of the argument in the function's
	// parameter list.
	ArgNo int
}

DIParameterVariable holds the values for creating parameter variable debug metadata.
type DIPointerType ¶

type DIPointerType struct {
	Pointee      Metadata
	SizeInBits   uint64
	AlignInBits  uint32 // optional
	AddressSpace uint32
	Name         string // optional
}

DIPointerType holds the values for creating pointer type debug metadata.
type DIReplaceableCompositeType ¶

type DIReplaceableCompositeType struct {
	Tag         dwarf.Tag
	Name        string
	File        Metadata
	Line        int
	RuntimeLang int
	SizeInBits  uint64
	AlignInBits uint32
	Flags       int
	UniqueID    string
}

DIReplaceableCompositeType holds the values for creating replaceable composite type debug metadata.
type DIStructType ¶

type DIStructType struct {
	Name         string
	File         Metadata
	Line         int
	SizeInBits   uint64
	AlignInBits  uint32
	Flags        int
	DerivedFrom  Metadata
	Elements     []Metadata
	VTableHolder Metadata // optional
	UniqueID     string
}

DIStructType holds the values for creating struct type debug metadata.
type DISubrange ¶

type DISubrange struct {
	Lo    int64
	Count int64
}

DISubrange describes an integer value range.
type DISubroutineType ¶

type DISubroutineType struct {
	// File is the file in which the subroutine type is defined.
	File Metadata

	// Parameters contains the subroutine parameter types,
	// including the return type at the 0th index.
	Parameters []Metadata

	Flags int
}

DISubroutineType holds the values for creating subroutine type debug metadata.
type DITypedef ¶

type DITypedef struct {
	Type        Metadata
	Name        string
	File        Metadata
	Line        int
	Context     Metadata
	AlignInBits uint32
}

DITypedef holds the values for creating typedef type debug metadata.
type DebugLoc ¶

type DebugLoc struct {
	Line, Col uint
	Scope     Metadata
	InlinedAt Metadata
}

Metadata
type DwarfLang ¶

type DwarfLang uint32

const (
	// http://dwarfstd.org/ShowIssue.php?issue=101014.1&type=open
	DW_LANG_Go DwarfLang = 0x0016
)

type DwarfTag ¶

type DwarfTag uint32

const (
	DW_TAG_lexical_block   DwarfTag = 0x0b
	DW_TAG_compile_unit    DwarfTag = 0x11
	DW_TAG_variable        DwarfTag = 0x34
	DW_TAG_base_type       DwarfTag = 0x24
	DW_TAG_pointer_type    DwarfTag = 0x0F
	DW_TAG_structure_type  DwarfTag = 0x13
	DW_TAG_subroutine_type DwarfTag = 0x15
	DW_TAG_file_type       DwarfTag = 0x29
	DW_TAG_subprogram      DwarfTag = 0x2E
	DW_TAG_auto_variable   DwarfTag = 0x100
	DW_TAG_arg_variable    DwarfTag = 0x101
)

type DwarfTypeEncoding ¶

type DwarfTypeEncoding uint32

const (
	DW_ATE_address         DwarfTypeEncoding = 0x01
	DW_ATE_boolean         DwarfTypeEncoding = 0x02
	DW_ATE_complex_float   DwarfTypeEncoding = 0x03
	DW_ATE_float           DwarfTypeEncoding = 0x04
	DW_ATE_signed          DwarfTypeEncoding = 0x05
	DW_ATE_signed_char     DwarfTypeEncoding = 0x06
	DW_ATE_unsigned        DwarfTypeEncoding = 0x07
	DW_ATE_unsigned_char   DwarfTypeEncoding = 0x08
	DW_ATE_imaginary_float DwarfTypeEncoding = 0x09
	DW_ATE_packed_decimal  DwarfTypeEncoding = 0x0a
	DW_ATE_numeric_string  DwarfTypeEncoding = 0x0b
	DW_ATE_edited          DwarfTypeEncoding = 0x0c
	DW_ATE_signed_fixed    DwarfTypeEncoding = 0x0d
	DW_ATE_unsigned_fixed  DwarfTypeEncoding = 0x0e
	DW_ATE_decimal_float   DwarfTypeEncoding = 0x0f
	DW_ATE_UTF             DwarfTypeEncoding = 0x10
	DW_ATE_lo_user         DwarfTypeEncoding = 0x80
	DW_ATE_hi_user         DwarfTypeEncoding = 0xff
)

type ExecutionEngine ¶

type ExecutionEngine struct {
	C C.LLVMExecutionEngineRef
}

func NewExecutionEngine ¶

func NewExecutionEngine(m Module) (ee ExecutionEngine, err error)

func NewInterpreter ¶

func NewInterpreter(m Module) (ee ExecutionEngine, err error)

func NewMCJITCompiler ¶

func NewMCJITCompiler(m Module, options MCJITCompilerOptions) (ee ExecutionEngine, err error)

func (ExecutionEngine) AddGlobalMapping ¶

func (ee ExecutionEngine) AddGlobalMapping(global Value, addr unsafe.Pointer)

func (ExecutionEngine) AddModule ¶

func (ee ExecutionEngine) AddModule(m Module)

func (ExecutionEngine) Dispose ¶

func (ee ExecutionEngine) Dispose()

func (ExecutionEngine) FindFunction ¶

func (ee ExecutionEngine) FindFunction(name string) (f Value)

func (ExecutionEngine) FreeMachineCodeForFunction ¶

func (ee ExecutionEngine) FreeMachineCodeForFunction(f Value)

func (ExecutionEngine) PointerToGlobal ¶

func (ee ExecutionEngine) PointerToGlobal(global Value) unsafe.Pointer

func (ExecutionEngine) RecompileAndRelinkFunction ¶

func (ee ExecutionEngine) RecompileAndRelinkFunction(f Value) unsafe.Pointer

func (ExecutionEngine) RemoveModule ¶

func (ee ExecutionEngine) RemoveModule(m Module)

func (ExecutionEngine) RunFunction ¶

func (ee ExecutionEngine) RunFunction(f Value, args []GenericValue) (g GenericValue)

func (ExecutionEngine) RunStaticConstructors ¶

func (ee ExecutionEngine) RunStaticConstructors()

func (ExecutionEngine) RunStaticDestructors ¶

func (ee ExecutionEngine) RunStaticDestructors()

func (ExecutionEngine) TargetData ¶

func (ee ExecutionEngine) TargetData() (td TargetData)

type FloatPredicate ¶

type FloatPredicate C.LLVMRealPredicate

const (
	FloatPredicateFalse FloatPredicate = C.LLVMRealPredicateFalse
	FloatOEQ            FloatPredicate = C.LLVMRealOEQ
	FloatOGT            FloatPredicate = C.LLVMRealOGT
	FloatOGE            FloatPredicate = C.LLVMRealOGE
	FloatOLT            FloatPredicate = C.LLVMRealOLT
	FloatOLE            FloatPredicate = C.LLVMRealOLE
	FloatONE            FloatPredicate = C.LLVMRealONE
	FloatORD            FloatPredicate = C.LLVMRealORD
	FloatUNO            FloatPredicate = C.LLVMRealUNO
	FloatUEQ            FloatPredicate = C.LLVMRealUEQ
	FloatUGT            FloatPredicate = C.LLVMRealUGT
	FloatUGE            FloatPredicate = C.LLVMRealUGE
	FloatULT            FloatPredicate = C.LLVMRealULT
	FloatULE            FloatPredicate = C.LLVMRealULE
	FloatUNE            FloatPredicate = C.LLVMRealUNE
	FloatPredicateTrue  FloatPredicate = C.LLVMRealPredicateTrue
)

type GenericValue ¶

type GenericValue struct {
	C C.LLVMGenericValueRef
}

func NewGenericValueFromFloat ¶

func NewGenericValueFromFloat(t Type, n float64) (g GenericValue)

func NewGenericValueFromInt ¶

func NewGenericValueFromInt(t Type, n uint64, signed bool) (g GenericValue)

func NewGenericValueFromPointer ¶

func NewGenericValueFromPointer(p unsafe.Pointer) (g GenericValue)

func (GenericValue) Dispose ¶

func (g GenericValue) Dispose()

func (GenericValue) Float ¶

func (g GenericValue) Float(t Type) float64

func (GenericValue) Int ¶

func (g GenericValue) Int(signed bool) uint64

func (GenericValue) IntWidth ¶

func (g GenericValue) IntWidth() int

func (GenericValue) Pointer ¶

func (g GenericValue) Pointer() unsafe.Pointer

type InlineAsmDialect ¶

type InlineAsmDialect C.LLVMInlineAsmDialect

const (
	InlineAsmDialectATT   InlineAsmDialect = C.LLVMInlineAsmDialectATT
	InlineAsmDialectIntel InlineAsmDialect = C.LLVMInlineAsmDialectIntel
)

type IntPredicate ¶

type IntPredicate C.LLVMIntPredicate

const (
	IntEQ  IntPredicate = C.LLVMIntEQ
	IntNE  IntPredicate = C.LLVMIntNE
	IntUGT IntPredicate = C.LLVMIntUGT
	IntUGE IntPredicate = C.LLVMIntUGE
	IntULT IntPredicate = C.LLVMIntULT
	IntULE IntPredicate = C.LLVMIntULE
	IntSGT IntPredicate = C.LLVMIntSGT
	IntSGE IntPredicate = C.LLVMIntSGE
	IntSLT IntPredicate = C.LLVMIntSLT
	IntSLE IntPredicate = C.LLVMIntSLE
)

type LandingPadClause ¶

type LandingPadClause C.LLVMLandingPadClauseTy

const (
	LandingPadCatch  LandingPadClause = C.LLVMLandingPadCatch
	LandingPadFilter LandingPadClause = C.LLVMLandingPadFilter
)

type Linkage ¶

type Linkage C.LLVMLinkage

const (
	ExternalLinkage            Linkage = C.LLVMExternalLinkage
	AvailableExternallyLinkage Linkage = C.LLVMAvailableExternallyLinkage
	LinkOnceAnyLinkage         Linkage = C.LLVMLinkOnceAnyLinkage
	LinkOnceODRLinkage         Linkage = C.LLVMLinkOnceODRLinkage
	WeakAnyLinkage             Linkage = C.LLVMWeakAnyLinkage
	WeakODRLinkage             Linkage = C.LLVMWeakODRLinkage
	AppendingLinkage           Linkage = C.LLVMAppendingLinkage
	InternalLinkage            Linkage = C.LLVMInternalLinkage
	PrivateLinkage             Linkage = C.LLVMPrivateLinkage
	ExternalWeakLinkage        Linkage = C.LLVMExternalWeakLinkage
	CommonLinkage              Linkage = C.LLVMCommonLinkage
)

type MCJITCompilerOptions ¶

type MCJITCompilerOptions struct {
	C C.struct_LLVMMCJITCompilerOptions
}

func NewMCJITCompilerOptions ¶

func NewMCJITCompilerOptions() MCJITCompilerOptions

func (*MCJITCompilerOptions) SetMCJITCodeModel ¶

func (options *MCJITCompilerOptions) SetMCJITCodeModel(CodeModel CodeModel)

func (*MCJITCompilerOptions) SetMCJITEnableFastISel ¶

func (options *MCJITCompilerOptions) SetMCJITEnableFastISel(fastisel bool)

func (*MCJITCompilerOptions) SetMCJITNoFramePointerElim ¶

func (options *MCJITCompilerOptions) SetMCJITNoFramePointerElim(nfp bool)

func (*MCJITCompilerOptions) SetMCJITOptimizationLevel ¶

func (options *MCJITCompilerOptions) SetMCJITOptimizationLevel(level uint)

type MemoryBuffer ¶

type MemoryBuffer struct {
	C C.LLVMMemoryBufferRef
}

func NewMemoryBufferFromFile ¶

func NewMemoryBufferFromFile(path string) (b MemoryBuffer, err error)

func NewMemoryBufferFromStdin ¶

func NewMemoryBufferFromStdin() (b MemoryBuffer, err error)

func WriteBitcodeToMemoryBuffer ¶

func WriteBitcodeToMemoryBuffer(m Module) MemoryBuffer

func WriteThinLTOBitcodeToMemoryBuffer ¶

func WriteThinLTOBitcodeToMemoryBuffer(m Module) MemoryBuffer

func (MemoryBuffer) Bytes ¶

func (b MemoryBuffer) Bytes() []byte

func (MemoryBuffer) Dispose ¶

func (b MemoryBuffer) Dispose()

func (MemoryBuffer) IsNil ¶

func (c MemoryBuffer) IsNil() bool

type Metadata ¶

type Metadata struct {
	C C.LLVMMetadataRef
}

func (Metadata) FileDirectory ¶

func (md Metadata) FileDirectory() string

FileDirectory returns the directory of a DIFile metadata node, or the empty string if there is no directory information.
func (Metadata) FileFilename ¶

func (md Metadata) FileFilename() string

FileFilename returns the filename of a DIFile metadata node, or the empty string if there is no filename information.
func (Metadata) FileSource ¶

func (md Metadata) FileSource() string

FileSource returns the source of a DIFile metadata node.
func (Metadata) IsNil ¶

func (c Metadata) IsNil() bool

func (Metadata) Kind ¶

func (md Metadata) Kind() MetadataKind

Kind returns the metadata kind.
func (Metadata) LocationColumn ¶

func (md Metadata) LocationColumn() uint

LocationColumn returns the column (offset from the start of the line) of a DILocation.
func (Metadata) LocationInlinedAt ¶

func (md Metadata) LocationInlinedAt() Metadata

LocationInlinedAt return the "inline at" location associated with this debug location.
func (Metadata) LocationLine ¶

func (md Metadata) LocationLine() uint

LocationLine returns the line number of a DILocation.
func (Metadata) LocationScope ¶

func (md Metadata) LocationScope() Metadata

LocationScope returns the local scope associated with this debug location.
func (Metadata) ReplaceAllUsesWith ¶

func (md Metadata) ReplaceAllUsesWith(new Metadata)

func (Metadata) ScopeFile ¶

func (md Metadata) ScopeFile() Metadata

ScopeFile returns the file (DIFile) of a given scope.
func (Metadata) SubprogramLine ¶

func (md Metadata) SubprogramLine() uint

SubprogramLine returns the line number of a DISubprogram.
type MetadataKind ¶

type MetadataKind C.LLVMMetadataKind

type Module ¶

type Module struct {
	C C.LLVMModuleRef
}

func (Module) AddNamedMetadataOperand ¶

func (m Module) AddNamedMetadataOperand(name string, operand Metadata)

func (Module) Comdat ¶

func (m Module) Comdat(name string) (c Comdat)

Operations on comdat
func (Module) Context ¶

func (m Module) Context() (c Context)

func (Module) DataLayout ¶

func (m Module) DataLayout() string

Data layout. See Module::getDataLayout.
func (Module) Dispose ¶

func (m Module) Dispose()

See llvm::Module::~Module
func (Module) Dump ¶

func (m Module) Dump()

See Module::dump.
func (Module) FirstFunction ¶

func (m Module) FirstFunction() (v Value)

func (Module) FirstGlobal ¶

func (m Module) FirstGlobal() (v Value)

func (Module) GetTypeByName ¶

func (m Module) GetTypeByName(name string) (t Type)

func (Module) IsNil ¶

func (c Module) IsNil() bool

func (Module) LastFunction ¶

func (m Module) LastFunction() (v Value)

func (Module) LastGlobal ¶

func (m Module) LastGlobal() (v Value)

func (Module) NamedFunction ¶

func (m Module) NamedFunction(name string) (v Value)

func (Module) NamedGlobal ¶

func (m Module) NamedGlobal(name string) (v Value)

func (Module) RunPasses ¶

func (mod Module) RunPasses(passes string, tm TargetMachine, options PassBuilderOptions) error

RunPasses runs the specified optimization passes on the functions in the module. `passes` is a comma separated list of pass names in the same format as llvm's `opt -passes=...` command. Running `opt -print-passes` can list the available passes.

Some notable passes include:

default<O0>  -- run the default -O0 passes
default<O1>  -- run the default -O1 passes
default<O2>  -- run the default -O2 passes
default<O3>  -- run the default -O3 passes
default<Os>  -- run the default -Os passes, like -O2 but size conscious
default<Oz>  -- run the default -Oz passes, optimizing for size above all else

func (Module) SetDataLayout ¶

func (m Module) SetDataLayout(layout string)

func (Module) SetInlineAsm ¶

func (m Module) SetInlineAsm(asm string)

See Module::setModuleInlineAsm.
func (Module) SetTarget ¶

func (m Module) SetTarget(target string)

func (Module) String ¶

func (m Module) String() string

func (Module) Target ¶

func (m Module) Target() string

Target triple. See Module::getTargetTriple.
type ModuleProvider ¶

type ModuleProvider struct {
	C C.LLVMModuleProviderRef
}

func NewModuleProviderForModule ¶

func NewModuleProviderForModule(m Module) (mp ModuleProvider)

Changes the type of M so it can be passed to FunctionPassManagers and the JIT. They take ModuleProviders for historical reasons.
func (ModuleProvider) Dispose ¶

func (mp ModuleProvider) Dispose()

Destroys the module M.
func (ModuleProvider) IsNil ¶

func (c ModuleProvider) IsNil() bool

type Opcode ¶

type Opcode C.LLVMOpcode

const (
	Ret         Opcode = C.LLVMRet
	Br          Opcode = C.LLVMBr
	Switch      Opcode = C.LLVMSwitch
	IndirectBr  Opcode = C.LLVMIndirectBr
	Invoke      Opcode = C.LLVMInvoke
	Unreachable Opcode = C.LLVMUnreachable

	// Standard Binary Operators
	Add  Opcode = C.LLVMAdd
	FAdd Opcode = C.LLVMFAdd
	Sub  Opcode = C.LLVMSub
	FSub Opcode = C.LLVMFSub
	Mul  Opcode = C.LLVMMul
	FMul Opcode = C.LLVMFMul
	UDiv Opcode = C.LLVMUDiv
	SDiv Opcode = C.LLVMSDiv
	FDiv Opcode = C.LLVMFDiv
	URem Opcode = C.LLVMURem
	SRem Opcode = C.LLVMSRem
	FRem Opcode = C.LLVMFRem

	// Logical Operators
	Shl  Opcode = C.LLVMShl
	LShr Opcode = C.LLVMLShr
	AShr Opcode = C.LLVMAShr
	And  Opcode = C.LLVMAnd
	Or   Opcode = C.LLVMOr
	Xor  Opcode = C.LLVMXor

	// Memory Operators
	Alloca        Opcode = C.LLVMAlloca
	Load          Opcode = C.LLVMLoad
	Store         Opcode = C.LLVMStore
	GetElementPtr Opcode = C.LLVMGetElementPtr

	// Cast Operators
	Trunc    Opcode = C.LLVMTrunc
	ZExt     Opcode = C.LLVMZExt
	SExt     Opcode = C.LLVMSExt
	FPToUI   Opcode = C.LLVMFPToUI
	FPToSI   Opcode = C.LLVMFPToSI
	UIToFP   Opcode = C.LLVMUIToFP
	SIToFP   Opcode = C.LLVMSIToFP
	FPTrunc  Opcode = C.LLVMFPTrunc
	FPExt    Opcode = C.LLVMFPExt
	PtrToInt Opcode = C.LLVMPtrToInt
	IntToPtr Opcode = C.LLVMIntToPtr
	BitCast  Opcode = C.LLVMBitCast

	// Other Operators
	ICmp   Opcode = C.LLVMICmp
	FCmp   Opcode = C.LLVMFCmp
	PHI    Opcode = C.LLVMPHI
	Call   Opcode = C.LLVMCall
	Select Opcode = C.LLVMSelect
	// UserOp1
	// UserOp2
	VAArg          Opcode = C.LLVMVAArg
	ExtractElement Opcode = C.LLVMExtractElement
	InsertElement  Opcode = C.LLVMInsertElement
	ShuffleVector  Opcode = C.LLVMShuffleVector
	ExtractValue   Opcode = C.LLVMExtractValue
	InsertValue    Opcode = C.LLVMInsertValue

	// Exception Handling Operators
	Resume      Opcode = C.LLVMResume
	LandingPad  Opcode = C.LLVMLandingPad
	CleanupRet  Opcode = C.LLVMCleanupRet
	CatchRet    Opcode = C.LLVMCatchRet
	CatchPad    Opcode = C.LLVMCatchPad
	CleanupPad  Opcode = C.LLVMCleanupPad
	CatchSwitch Opcode = C.LLVMCatchSwitch
)

type PassBuilderOptions ¶

type PassBuilderOptions struct {
	C C.LLVMPassBuilderOptionsRef
}

PassBuilderOptions allows specifying several options for the PassBuilder.
func NewPassBuilderOptions ¶

func NewPassBuilderOptions() (pbo PassBuilderOptions)

NewPassBuilderOptions creates a PassBuilderOptions which can be used to specify pass options for RunPasses.
func (PassBuilderOptions) Dispose ¶

func (pbo PassBuilderOptions) Dispose()

Dispose of the memory allocated for the PassBuilderOptions.
func (PassBuilderOptions) SetCallGraphProfile ¶

func (pbo PassBuilderOptions) SetCallGraphProfile(cgProfile bool)

SetCallGraphProfile toggles whether call graph profiling should be used.
func (PassBuilderOptions) SetDebugLogging ¶

func (pbo PassBuilderOptions) SetDebugLogging(debugLogging bool)

SetDebugLogging toggles debug logging for the PassBuilder.
func (PassBuilderOptions) SetForgetAllSCEVInLoopUnroll ¶

func (pbo PassBuilderOptions) SetForgetAllSCEVInLoopUnroll(forgetSCEV bool)

SetForgetAllSCEVInLoopUnroll toggles forgetting all SCEV (Scalar Evolution) information in loop unrolling. Scalar Evolution is a pass that analyses the how scalars evolve over iterations of a loop in order to optimize the loop better. Forgetting this information can be useful in some cases.
func (PassBuilderOptions) SetLicmMssaNoAccForPromotionCap ¶

func (pbo PassBuilderOptions) SetLicmMssaNoAccForPromotionCap(promotionCap uint)

SetLicmMssaNoAccForPromotionCap sets a tuning option to cap the number of promotions to scalars in Loop Invariant Code Motion with MemorySSA, if the number of accesses is too large. See [llvm::PipelineTuningOptions::LicmMssaNoAccForPromotionCap].
func (PassBuilderOptions) SetLicmMssaOptCap ¶

func (pbo PassBuilderOptions) SetLicmMssaOptCap(optCap uint)

SetLicmMssaOptCap sets a tuning option to cap the number of calls to retrieve clobbering accesses in MemorySSA, in Loop Invariant Code Motion optimization. See [llvm::PipelineTuningOptions::LicmMssaOptCap].
func (PassBuilderOptions) SetLoopInterleaving ¶

func (pbo PassBuilderOptions) SetLoopInterleaving(loopInterleaving bool)

SetLoopInterleaving toggles loop interleaving, which is part of loop vectorization.
func (PassBuilderOptions) SetLoopUnrolling ¶

func (pbo PassBuilderOptions) SetLoopUnrolling(loopUnrolling bool)

SetLoopUnrolling toggles loop unrolling.
func (PassBuilderOptions) SetLoopVectorization ¶

func (pbo PassBuilderOptions) SetLoopVectorization(loopVectorization bool)

SetLoopVectorization toggles loop vectorization.
func (PassBuilderOptions) SetMergeFunctions ¶

func (pbo PassBuilderOptions) SetMergeFunctions(mergeFuncs bool)

SetMergeFunctions toggles finding functions which will generate identical machine code by considering all pointer types to be equivalent. Once identified, they will be folded by replacing a call to one with a call to a bitcast of the other.
func (PassBuilderOptions) SetSLPVectorization ¶

func (pbo PassBuilderOptions) SetSLPVectorization(slpVectorization bool)

SetSLPVectorization toggles Super-Word Level Parallelism vectorization, whose goal is to combine multiple similar independent instructions into a vector instruction.
func (PassBuilderOptions) SetVerifyEach ¶

func (pbo PassBuilderOptions) SetVerifyEach(verifyEach bool)

SetVerifyEach toggles adding a VerifierPass to the PassBuilder, ensuring all functions inside the module are valid. Useful for debugging, but adds a significant amount of overhead.
type PassManager ¶

type PassManager struct {
	C C.LLVMPassManagerRef
}

func NewFunctionPassManagerForModule ¶

func NewFunctionPassManagerForModule(m Module) (pm PassManager)

Constructs a new function-by-function pass pipeline over the module provider. It does not take ownership of the module provider. This type of pipeline is suitable for code generation and JIT compilation tasks. See llvm::FunctionPassManager::FunctionPassManager.
func NewPassManager ¶

func NewPassManager() (pm PassManager)

Constructs a new whole-module pass pipeline. This type of pipeline is suitable for link-time optimization and whole-module transformations. See llvm::PassManager::PassManager.
func (PassManager) Dispose ¶

func (pm PassManager) Dispose()

Frees the memory of a pass pipeline. For function pipelines, does not free the module provider. See llvm::PassManagerBase::~PassManagerBase.
func (PassManager) FinalizeFunc ¶

func (pm PassManager) FinalizeFunc() bool

Finalizes all of the function passes scheduled in the function pass manager. Returns 1 if any of the passes modified the module, 0 otherwise. See llvm::FunctionPassManager::doFinalization.
func (PassManager) InitializeFunc ¶

func (pm PassManager) InitializeFunc() bool

Initializes all of the function passes scheduled in the function pass manager. Returns 1 if any of the passes modified the module, 0 otherwise. See llvm::FunctionPassManager::doInitialization.
func (PassManager) IsNil ¶

func (c PassManager) IsNil() bool

func (PassManager) Run ¶

func (pm PassManager) Run(m Module) bool

Initializes, executes on the provided module, and finalizes all of the passes scheduled in the pass manager. Returns 1 if any of the passes modified the module, 0 otherwise. See llvm::PassManager::run(Module&).
func (PassManager) RunFunc ¶

func (pm PassManager) RunFunc(f Value) bool

Executes all of the function passes scheduled in the function pass manager on the provided function. Returns 1 if any of the passes modified the function, false otherwise. See llvm::FunctionPassManager::run(Function&).
type RelocMode ¶

type RelocMode C.LLVMRelocMode

const (
	RelocDefault      RelocMode = C.LLVMRelocDefault
	RelocStatic       RelocMode = C.LLVMRelocStatic
	RelocPIC          RelocMode = C.LLVMRelocPIC
	RelocDynamicNoPic RelocMode = C.LLVMRelocDynamicNoPic
)

type Target ¶

type Target struct {
	C C.LLVMTargetRef
}

func FirstTarget ¶

func FirstTarget() Target

func GetTargetFromTriple ¶

func GetTargetFromTriple(triple string) (t Target, err error)

func (Target) CreateTargetMachine ¶

func (t Target) CreateTargetMachine(Triple string, CPU string, Features string,
	Level CodeGenOptLevel, Reloc RelocMode,
	CodeModel CodeModel) (tm TargetMachine)

CreateTargetMachine creates a new TargetMachine.
func (Target) Description ¶

func (t Target) Description() string

func (Target) Name ¶

func (t Target) Name() string

func (Target) NextTarget ¶

func (t Target) NextTarget() Target

type TargetData ¶

type TargetData struct {
	C C.LLVMTargetDataRef
}

func NewTargetData ¶

func NewTargetData(rep string) (td TargetData)

Creates target data from a target layout string. See the constructor llvm::TargetData::TargetData.
func (TargetData) ABITypeAlignment ¶

func (td TargetData) ABITypeAlignment(t Type) int

Computes the ABI alignment of a type in bytes for a target. See the method llvm::TargetData::getABITypeAlignment.
func (TargetData) ByteOrder ¶

func (td TargetData) ByteOrder() ByteOrdering

Returns the byte order of a target, either BigEndian or LittleEndian. See the method llvm::TargetData::isLittleEndian.
func (TargetData) CallFrameTypeAlignment ¶

func (td TargetData) CallFrameTypeAlignment(t Type) int

Computes the call frame alignment of a type in bytes for a target. See the method llvm::TargetData::getCallFrameTypeAlignment.
func (TargetData) Dispose ¶

func (td TargetData) Dispose()

Deallocates a TargetData. See the destructor llvm::TargetData::~TargetData.
func (TargetData) ElementContainingOffset ¶

func (td TargetData) ElementContainingOffset(t Type, offset uint64) int

Computes the structure element that contains the byte offset for a target. See the method llvm::StructLayout::getElementContainingOffset.
func (TargetData) ElementOffset ¶

func (td TargetData) ElementOffset(t Type, element int) uint64

Computes the byte offset of the indexed struct element for a target. See the method llvm::StructLayout::getElementOffset.
func (TargetData) IntPtrType ¶

func (td TargetData) IntPtrType() (t Type)

Returns the integer type that is the same size as a pointer on a target. See the method llvm::TargetData::getIntPtrType.
func (TargetData) PointerSize ¶

func (td TargetData) PointerSize() int

Returns the pointer size in bytes for a target. See the method llvm::TargetData::getPointerSize.
func (TargetData) PrefTypeAlignment ¶

func (td TargetData) PrefTypeAlignment(t Type) int

Computes the preferred alignment of a type in bytes for a target. See the method llvm::TargetData::getPrefTypeAlignment.
func (TargetData) PreferredAlignment ¶

func (td TargetData) PreferredAlignment(g Value) int

Computes the preferred alignment of a global variable in bytes for a target. See the method llvm::TargetData::getPreferredAlignment.
func (TargetData) String ¶

func (td TargetData) String() (s string)

Converts target data to a target layout string. The string must be disposed with LLVMDisposeMessage. See the constructor llvm::TargetData::TargetData.
func (TargetData) TypeAllocSize ¶

func (td TargetData) TypeAllocSize(t Type) uint64

Computes the ABI size of a type in bytes for a target. See the method llvm::TargetData::getTypeAllocSize.
func (TargetData) TypeSizeInBits ¶

func (td TargetData) TypeSizeInBits(t Type) uint64

Computes the size of a type in bytes for a target. See the method llvm::TargetData::getTypeSizeInBits.
func (TargetData) TypeStoreSize ¶

func (td TargetData) TypeStoreSize(t Type) uint64

Computes the storage size of a type in bytes for a target. See the method llvm::TargetData::getTypeStoreSize.
type TargetMachine ¶

type TargetMachine struct {
	C C.LLVMTargetMachineRef
}

func (TargetMachine) AddAnalysisPasses ¶

func (tm TargetMachine) AddAnalysisPasses(pm PassManager)

func (TargetMachine) CreateTargetData ¶

func (tm TargetMachine) CreateTargetData() TargetData

CreateTargetData returns a new TargetData describing the TargetMachine's data layout. The returned TargetData is owned by the caller, who is responsible for disposing of it by calling the TargetData.Dispose method.
func (TargetMachine) Dispose ¶

func (tm TargetMachine) Dispose()

Dispose releases resources related to the TargetMachine.
func (TargetMachine) EmitToMemoryBuffer ¶

func (tm TargetMachine) EmitToMemoryBuffer(m Module, ft CodeGenFileType) (MemoryBuffer, error)

func (TargetMachine) Triple ¶

func (tm TargetMachine) Triple() string

Triple returns the triple describing the machine (arch-vendor-os).
type Type ¶

type Type struct {
	C C.LLVMTypeRef
}

func ArrayType ¶

func ArrayType(elementType Type, elementCount int) (t Type)

func FunctionType ¶

func FunctionType(returnType Type, paramTypes []Type, isVarArg bool) (t Type)

Operations on function types
func PointerType ¶

func PointerType(elementType Type, addressSpace int) (t Type)

func StructType ¶

func StructType(elementTypes []Type, packed bool) (t Type)

func VectorType ¶

func VectorType(elementType Type, elementCount int) (t Type)

func (Type) ArrayLength ¶

func (t Type) ArrayLength() int

func (Type) Context ¶

func (t Type) Context() (c Context)

See llvm::LLVMType::getContext.
func (Type) ElementType ¶

func (t Type) ElementType() (rt Type)

func (Type) IntTypeWidth ¶

func (t Type) IntTypeWidth() int

func (Type) IsFunctionVarArg ¶

func (t Type) IsFunctionVarArg() bool

func (Type) IsNil ¶

func (c Type) IsNil() bool

func (Type) IsStructPacked ¶

func (t Type) IsStructPacked() bool

func (Type) ParamTypes ¶

func (t Type) ParamTypes() []Type

func (Type) ParamTypesCount ¶

func (t Type) ParamTypesCount() int

func (Type) PointerAddressSpace ¶

func (t Type) PointerAddressSpace() int

func (Type) ReturnType ¶

func (t Type) ReturnType() (rt Type)

func (Type) String ¶

func (t Type) String() string

func (Type) StructElementTypes ¶

func (t Type) StructElementTypes() []Type

func (Type) StructElementTypesCount ¶

func (t Type) StructElementTypesCount() int

func (Type) StructName ¶

func (t Type) StructName() string

func (Type) StructSetBody ¶

func (t Type) StructSetBody(elementTypes []Type, packed bool)

func (Type) Subtypes ¶

func (t Type) Subtypes() (ret []Type)

Operations on array, pointer, and vector types (sequence types)
func (Type) TypeKind ¶

func (t Type) TypeKind() TypeKind

See llvm::LLVMTypeKind::getTypeID.
func (Type) VectorSize ¶

func (t Type) VectorSize() int

type TypeKind ¶

type TypeKind C.LLVMTypeKind

const (
	VoidTypeKind      TypeKind = C.LLVMVoidTypeKind
	FloatTypeKind     TypeKind = C.LLVMFloatTypeKind
	DoubleTypeKind    TypeKind = C.LLVMDoubleTypeKind
	X86_FP80TypeKind  TypeKind = C.LLVMX86_FP80TypeKind
	FP128TypeKind     TypeKind = C.LLVMFP128TypeKind
	PPC_FP128TypeKind TypeKind = C.LLVMPPC_FP128TypeKind
	LabelTypeKind     TypeKind = C.LLVMLabelTypeKind
	IntegerTypeKind   TypeKind = C.LLVMIntegerTypeKind
	FunctionTypeKind  TypeKind = C.LLVMFunctionTypeKind
	StructTypeKind    TypeKind = C.LLVMStructTypeKind
	ArrayTypeKind     TypeKind = C.LLVMArrayTypeKind
	PointerTypeKind   TypeKind = C.LLVMPointerTypeKind
	VectorTypeKind    TypeKind = C.LLVMVectorTypeKind
	MetadataTypeKind  TypeKind = C.LLVMMetadataTypeKind
	TokenTypeKind     TypeKind = C.LLVMTokenTypeKind
)

func (TypeKind) String ¶

func (t TypeKind) String() string

type Use ¶

type Use struct {
	C C.LLVMUseRef
}

func (Use) IsNil ¶

func (c Use) IsNil() bool

func (Use) NextUse ¶

func (u Use) NextUse() (ru Use)

func (Use) UsedValue ¶

func (u Use) UsedValue() (v Value)

func (Use) User ¶

func (u Use) User() (v Value)

type Value ¶

type Value struct {
	C C.LLVMValueRef
}

func AddAlias ¶

func AddAlias(m Module, t Type, addressSpace int, aliasee Value, name string) (v Value)

Operations on aliases
func AddFunction ¶

func AddFunction(m Module, name string, ft Type) (v Value)

Operations on functions
func AddGlobal ¶

func AddGlobal(m Module, t Type, name string) (v Value)

Operations on global variables
func AddGlobalInAddressSpace ¶

func AddGlobalInAddressSpace(m Module, t Type, name string, addressSpace int) (v Value)

func AlignOf ¶

func AlignOf(t Type) (v Value)

func BlockAddress ¶

func BlockAddress(f Value, bb BasicBlock) (v Value)

func ConstAdd ¶

func ConstAdd(lhs, rhs Value) (v Value)

func ConstAllOnes ¶

func ConstAllOnes(t Type) (v Value)

func ConstArray ¶

func ConstArray(t Type, constVals []Value) (v Value)

func ConstBitCast ¶

func ConstBitCast(v Value, t Type) (rv Value)

func ConstExtractElement ¶

func ConstExtractElement(vec, i Value) (rv Value)

func ConstFloat ¶

func ConstFloat(t Type, n float64) (v Value)

func ConstFloatFromString ¶

func ConstFloatFromString(t Type, str string) (v Value)

func ConstGEP ¶

func ConstGEP(t Type, v Value, indices []Value) (rv Value)

func ConstInBoundsGEP ¶

func ConstInBoundsGEP(t Type, v Value, indices []Value) (rv Value)

func ConstInsertElement ¶

func ConstInsertElement(vec, elem, i Value) (rv Value)

func ConstInt ¶

func ConstInt(t Type, n uint64, signExtend bool) (v Value)

Operations on scalar constants
func ConstIntFromString ¶

func ConstIntFromString(t Type, str string, radix int) (v Value)

func ConstIntToPtr ¶

func ConstIntToPtr(v Value, t Type) (rv Value)

func ConstMul ¶

func ConstMul(lhs, rhs Value) (v Value)

func ConstNSWAdd ¶

func ConstNSWAdd(lhs, rhs Value) (v Value)

func ConstNSWMul ¶

func ConstNSWMul(lhs, rhs Value) (v Value)

func ConstNSWNeg ¶

func ConstNSWNeg(v Value) (rv Value)

func ConstNSWSub ¶

func ConstNSWSub(lhs, rhs Value) (v Value)

func ConstNUWAdd ¶

func ConstNUWAdd(lhs, rhs Value) (v Value)

func ConstNUWMul ¶

func ConstNUWMul(lhs, rhs Value) (v Value)

func ConstNUWSub ¶

func ConstNUWSub(lhs, rhs Value) (v Value)

func ConstNamedStruct ¶

func ConstNamedStruct(t Type, constVals []Value) (v Value)

func ConstNeg ¶

func ConstNeg(v Value) (rv Value)

func ConstNot ¶

func ConstNot(v Value) (rv Value)

func ConstNull ¶

func ConstNull(t Type) (v Value)

Operations on constants of any type
func ConstPointerCast ¶

func ConstPointerCast(v Value, t Type) (rv Value)

func ConstPointerNull ¶

func ConstPointerNull(t Type) (v Value)

func ConstPtrToInt ¶

func ConstPtrToInt(v Value, t Type) (rv Value)

func ConstShuffleVector ¶

func ConstShuffleVector(veca, vecb, mask Value) (rv Value)

func ConstString ¶

func ConstString(str string, addnull bool) (v Value)

func ConstStruct ¶

func ConstStruct(constVals []Value, packed bool) (v Value)

func ConstSub ¶

func ConstSub(lhs, rhs Value) (v Value)

func ConstTrunc ¶

func ConstTrunc(v Value, t Type) (rv Value)

func ConstTruncOrBitCast ¶

func ConstTruncOrBitCast(v Value, t Type) (rv Value)

func ConstVector ¶

func ConstVector(scalarConstVals []Value, packed bool) (v Value)

func ConstXor ¶

func ConstXor(lhs, rhs Value) (v Value)

func InlineAsm ¶

func InlineAsm(t Type, asmString, constraints string, hasSideEffects, isAlignStack bool, dialect InlineAsmDialect, canThrow bool) (rv Value)

Operations on inline assembly
func NextFunction ¶

func NextFunction(v Value) (rv Value)

func NextGlobal ¶

func NextGlobal(v Value) (rv Value)

func NextInstruction ¶

func NextInstruction(v Value) (rv Value)

func NextParam ¶

func NextParam(v Value) (rv Value)

func PrevFunction ¶

func PrevFunction(v Value) (rv Value)

func PrevGlobal ¶

func PrevGlobal(v Value) (rv Value)

func PrevInstruction ¶

func PrevInstruction(v Value) (rv Value)

func PrevParam ¶

func PrevParam(v Value) (rv Value)

func SizeOf ¶

func SizeOf(t Type) (v Value)

func Undef ¶

func Undef(t Type) (v Value)

func (Value) AddAttributeAtIndex ¶

func (v Value) AddAttributeAtIndex(i int, a Attribute)

func (Value) AddCallSiteAttribute ¶

func (v Value) AddCallSiteAttribute(i int, a Attribute)

func (Value) AddCase ¶

func (v Value) AddCase(on Value, dest BasicBlock)

Add a case to the switch instruction
func (Value) AddClause ¶

func (l Value) AddClause(v Value)

func (Value) AddDest ¶

func (v Value) AddDest(dest BasicBlock)

Add a destination to the indirectbr instruction
func (Value) AddFunctionAttr ¶

func (v Value) AddFunctionAttr(a Attribute)

func (Value) AddHandler ¶

func (v Value) AddHandler(bb BasicBlock)

Add a destination to the catchswitch instruction.
func (Value) AddIncoming ¶

func (v Value) AddIncoming(vals []Value, blocks []BasicBlock)

Operations on phi nodes
func (Value) AddMetadata ¶

func (v Value) AddMetadata(kind int, md Metadata)

AddMetadata adds a metadata entry of the given kind to a global object.
func (Value) AddTargetDependentFunctionAttr ¶

func (v Value) AddTargetDependentFunctionAttr(attr, value string)

func (Value) Alignment ¶

func (v Value) Alignment() int

func (Value) AllocatedType ¶

func (v Value) AllocatedType() (t Type)

Operations on allocas
func (Value) AsBasicBlock ¶

func (v Value) AsBasicBlock() (bb BasicBlock)

func (Value) BasicBlocks ¶

func (v Value) BasicBlocks() []BasicBlock

func (Value) BasicBlocksCount ¶

func (v Value) BasicBlocksCount() int

func (Value) CalledFunctionType ¶

func (v Value) CalledFunctionType() (t Type)

func (Value) CalledValue ¶

func (v Value) CalledValue() (rv Value)

func (Value) CmpXchgFailureOrdering ¶

func (v Value) CmpXchgFailureOrdering() AtomicOrdering

func (Value) CmpXchgSuccessOrdering ¶

func (v Value) CmpXchgSuccessOrdering() AtomicOrdering

func (Value) Comdat ¶

func (v Value) Comdat() (c Comdat)

func (Value) ConstGetAsString ¶

func (v Value) ConstGetAsString() string

ConstGetAsString will return the string contained in a constant.
func (Value) ConstantAsMetadata ¶

func (v Value) ConstantAsMetadata() (md Metadata)

func (Value) DoubleValue ¶

func (v Value) DoubleValue() (result float64, inexact bool)

func (Value) Dump ¶

func (v Value) Dump()

func (Value) EntryBasicBlock ¶

func (v Value) EntryBasicBlock() (bb BasicBlock)

func (Value) EraseFromParentAsFunction ¶

func (v Value) EraseFromParentAsFunction()

func (Value) EraseFromParentAsGlobal ¶

func (v Value) EraseFromParentAsGlobal()

func (Value) EraseFromParentAsInstruction ¶

func (v Value) EraseFromParentAsInstruction()

Operations on instructions
func (Value) FirstBasicBlock ¶

func (v Value) FirstBasicBlock() (bb BasicBlock)

func (Value) FirstParam ¶

func (v Value) FirstParam() (rv Value)

func (Value) FirstUse ¶

func (v Value) FirstUse() (u Use)

Operations on Uses
func (Value) FloatPredicate ¶

func (v Value) FloatPredicate() FloatPredicate

func (Value) FunctionCallConv ¶

func (v Value) FunctionCallConv() CallConv

func (Value) GC ¶

func (v Value) GC() string

func (Value) GEPSourceElementType ¶

func (v Value) GEPSourceElementType() (t Type)

Operations on GEPs
func (Value) GetCallSiteEnumAttribute ¶

func (v Value) GetCallSiteEnumAttribute(i int, kind uint) (a Attribute)

func (Value) GetCallSiteStringAttribute ¶

func (v Value) GetCallSiteStringAttribute(i int, kind string) (a Attribute)

func (Value) GetEnumAttributeAtIndex ¶

func (v Value) GetEnumAttributeAtIndex(i int, kind uint) (a Attribute)

func (Value) GetEnumFunctionAttribute ¶

func (v Value) GetEnumFunctionAttribute(kind uint) Attribute

func (Value) GetHandlers ¶

func (v Value) GetHandlers() []BasicBlock

Obtain the basic blocks acting as handlers for a catchswitch instruction.
func (Value) GetParentCatchSwitch ¶

func (v Value) GetParentCatchSwitch() (rv Value)

Get the parent catchswitch instruction of a catchpad instruction.

This only works on catchpad instructions.
func (Value) GetStringAttributeAtIndex ¶

func (v Value) GetStringAttributeAtIndex(i int, kind string) (a Attribute)

func (Value) GlobalParent ¶

func (v Value) GlobalParent() (m Module)

Operations on global variables, functions, and aliases (globals)
func (Value) GlobalValueType ¶

func (v Value) GlobalValueType() (t Type)

func (Value) HasMetadata ¶

func (v Value) HasMetadata() bool

func (Value) IncomingBlock ¶

func (v Value) IncomingBlock(i int) (bb BasicBlock)

func (Value) IncomingCount ¶

func (v Value) IncomingCount() int

func (Value) IncomingValue ¶

func (v Value) IncomingValue(i int) (rv Value)

func (Value) Indices ¶

func (v Value) Indices() []uint32

Operations on aggregates
func (Value) Initializer ¶

func (v Value) Initializer() (rv Value)

func (Value) InstructionCallConv ¶

func (v Value) InstructionCallConv() CallConv

func (Value) InstructionDebugLoc ¶

func (v Value) InstructionDebugLoc() (md Metadata)

func (Value) InstructionOpcode ¶

func (v Value) InstructionOpcode() Opcode

func (Value) InstructionParent ¶

func (v Value) InstructionParent() (bb BasicBlock)

func (Value) InstructionSetDebugLoc ¶

func (v Value) InstructionSetDebugLoc(md Metadata)

func (Value) IntPredicate ¶

func (v Value) IntPredicate() IntPredicate

Operations on comparisons
func (Value) IntrinsicID ¶

func (v Value) IntrinsicID() int

func (Value) IsAAllocaInst ¶

func (v Value) IsAAllocaInst() (rv Value)

func (Value) IsAArgument ¶

func (v Value) IsAArgument() (rv Value)

Conversion functions. Return the input value if it is an instance of the specified class, otherwise NULL. See llvm::dyn_cast_or_null<>.
func (Value) IsABasicBlock ¶

func (v Value) IsABasicBlock() (rv Value)

func (Value) IsABinaryOperator ¶

func (v Value) IsABinaryOperator() (rv Value)

func (Value) IsABitCastInst ¶

func (v Value) IsABitCastInst() (rv Value)

func (Value) IsABranchInst ¶

func (v Value) IsABranchInst() (rv Value)

func (Value) IsACallInst ¶

func (v Value) IsACallInst() (rv Value)

func (Value) IsACastInst ¶

func (v Value) IsACastInst() (rv Value)

func (Value) IsACmpInst ¶

func (v Value) IsACmpInst() (rv Value)

func (Value) IsAConstant ¶

func (v Value) IsAConstant() (rv Value)

func (Value) IsAConstantAggregateZero ¶

func (v Value) IsAConstantAggregateZero() (rv Value)

func (Value) IsAConstantArray ¶

func (v Value) IsAConstantArray() (rv Value)

func (Value) IsAConstantExpr ¶

func (v Value) IsAConstantExpr() (rv Value)

func (Value) IsAConstantFP ¶

func (v Value) IsAConstantFP() (rv Value)

func (Value) IsAConstantInt ¶

func (v Value) IsAConstantInt() (rv Value)

func (Value) IsAConstantPointerNull ¶

func (v Value) IsAConstantPointerNull() (rv Value)

func (Value) IsAConstantStruct ¶

func (v Value) IsAConstantStruct() (rv Value)

func (Value) IsAConstantVector ¶

func (v Value) IsAConstantVector() (rv Value)

func (Value) IsADbgDeclareInst ¶

func (v Value) IsADbgDeclareInst() (rv Value)

func (Value) IsADbgInfoIntrinsic ¶

func (v Value) IsADbgInfoIntrinsic() (rv Value)

func (Value) IsAExtractElementInst ¶

func (v Value) IsAExtractElementInst() (rv Value)

func (Value) IsAExtractValueInst ¶

func (v Value) IsAExtractValueInst() (rv Value)

func (Value) IsAFCmpInst ¶

func (v Value) IsAFCmpInst() (rv Value)

func (Value) IsAFPExtInst ¶

func (v Value) IsAFPExtInst() (rv Value)

func (Value) IsAFPToSIInst ¶

func (v Value) IsAFPToSIInst() (rv Value)

func (Value) IsAFPToUIInst ¶

func (v Value) IsAFPToUIInst() (rv Value)

func (Value) IsAFPTruncInst ¶

func (v Value) IsAFPTruncInst() (rv Value)

func (Value) IsAFunction ¶

func (v Value) IsAFunction() (rv Value)

func (Value) IsAGetElementPtrInst ¶

func (v Value) IsAGetElementPtrInst() (rv Value)

func (Value) IsAGlobalAlias ¶

func (v Value) IsAGlobalAlias() (rv Value)

func (Value) IsAGlobalValue ¶

func (v Value) IsAGlobalValue() (rv Value)

func (Value) IsAGlobalVariable ¶

func (v Value) IsAGlobalVariable() (rv Value)

func (Value) IsAICmpInst ¶

func (v Value) IsAICmpInst() (rv Value)

func (Value) IsAInlineAsm ¶

func (v Value) IsAInlineAsm() (rv Value)

func (Value) IsAInsertElementInst ¶

func (v Value) IsAInsertElementInst() (rv Value)

func (Value) IsAInsertValueInst ¶

func (v Value) IsAInsertValueInst() (rv Value)

func (Value) IsAInstruction ¶

func (v Value) IsAInstruction() (rv Value)

func (Value) IsAIntToPtrInst ¶

func (v Value) IsAIntToPtrInst() (rv Value)

func (Value) IsAIntrinsicInst ¶

func (v Value) IsAIntrinsicInst() (rv Value)

func (Value) IsAInvokeInst ¶

func (v Value) IsAInvokeInst() (rv Value)

func (Value) IsALoadInst ¶

func (v Value) IsALoadInst() (rv Value)

func (Value) IsAMemCpyInst ¶

func (v Value) IsAMemCpyInst() (rv Value)

func (Value) IsAMemIntrinsic ¶

func (v Value) IsAMemIntrinsic() (rv Value)

func (Value) IsAMemMoveInst ¶

func (v Value) IsAMemMoveInst() (rv Value)

func (Value) IsAMemSetInst ¶

func (v Value) IsAMemSetInst() (rv Value)

func (Value) IsAPHINode ¶

func (v Value) IsAPHINode() (rv Value)

func (Value) IsAPtrToIntInst ¶

func (v Value) IsAPtrToIntInst() (rv Value)

func (Value) IsAReturnInst ¶

func (v Value) IsAReturnInst() (rv Value)

func (Value) IsASExtInst ¶

func (v Value) IsASExtInst() (rv Value)

func (Value) IsASIToFPInst ¶

func (v Value) IsASIToFPInst() (rv Value)

func (Value) IsASelectInst ¶

func (v Value) IsASelectInst() (rv Value)

func (Value) IsAShuffleVectorInst ¶

func (v Value) IsAShuffleVectorInst() (rv Value)

func (Value) IsAStoreInst ¶

func (v Value) IsAStoreInst() (rv Value)

func (Value) IsASwitchInst ¶

func (v Value) IsASwitchInst() (rv Value)

func (Value) IsATruncInst ¶

func (v Value) IsATruncInst() (rv Value)

func (Value) IsAUIToFPInst ¶

func (v Value) IsAUIToFPInst() (rv Value)

func (Value) IsAUnaryInstruction ¶

func (v Value) IsAUnaryInstruction() (rv Value)

func (Value) IsAUndefValue ¶

func (v Value) IsAUndefValue() (rv Value)

func (Value) IsAUnreachableInst ¶

func (v Value) IsAUnreachableInst() (rv Value)

func (Value) IsAUser ¶

func (v Value) IsAUser() (rv Value)

func (Value) IsAVAArgInst ¶

func (v Value) IsAVAArgInst() (rv Value)

func (Value) IsAZExtInst ¶

func (v Value) IsAZExtInst() (rv Value)

func (Value) IsAtomicSingleThread ¶

func (v Value) IsAtomicSingleThread() bool

func (Value) IsBasicBlock ¶

func (v Value) IsBasicBlock() bool

func (Value) IsConstant ¶

func (v Value) IsConstant() bool

func (Value) IsConstantString ¶

func (v Value) IsConstantString() bool

IsConstantString checks if the constant is an array of i8.
func (Value) IsDeclaration ¶

func (v Value) IsDeclaration() bool

func (Value) IsGlobalConstant ¶

func (v Value) IsGlobalConstant() bool

func (Value) IsNil ¶

func (c Value) IsNil() bool

func (Value) IsNull ¶

func (v Value) IsNull() bool

func (Value) IsTailCall ¶

func (v Value) IsTailCall() bool

Operations on call instructions (only)
func (Value) IsThreadLocal ¶

func (v Value) IsThreadLocal() bool

func (Value) IsUndef ¶

func (v Value) IsUndef() bool

func (Value) IsVolatile ¶

func (v Value) IsVolatile() bool

func (Value) LastBasicBlock ¶

func (v Value) LastBasicBlock() (bb BasicBlock)

func (Value) LastParam ¶

func (v Value) LastParam() (rv Value)

func (Value) Linkage ¶

func (v Value) Linkage() Linkage

func (Value) Metadata ¶

func (v Value) Metadata(kind int) (rv Value)

func (Value) Name ¶

func (v Value) Name() string

func (Value) Opcode ¶

func (v Value) Opcode() Opcode

Constant expressions
func (Value) Operand ¶

func (v Value) Operand(i int) (rv Value)

Operations on Users
func (Value) OperandsCount ¶

func (v Value) OperandsCount() int

func (Value) Ordering ¶

func (v Value) Ordering() AtomicOrdering

func (Value) Param ¶

func (v Value) Param(i int) (rv Value)

func (Value) ParamParent ¶

func (v Value) ParamParent() (rv Value)

func (Value) Params ¶

func (v Value) Params() []Value

func (Value) ParamsCount ¶

func (v Value) ParamsCount() int

Operations on parameters
func (Value) RemoveEnumAttributeAtIndex ¶

func (v Value) RemoveEnumAttributeAtIndex(i int, kind uint)

func (Value) RemoveEnumFunctionAttribute ¶

func (v Value) RemoveEnumFunctionAttribute(kind uint)

func (Value) RemoveFromParentAsInstruction ¶

func (v Value) RemoveFromParentAsInstruction()

func (Value) RemoveStringAttributeAtIndex ¶

func (v Value) RemoveStringAttributeAtIndex(i int, kind string)

func (Value) ReplaceAllUsesWith ¶

func (v Value) ReplaceAllUsesWith(nv Value)

func (Value) SExtValue ¶

func (v Value) SExtValue() int64

func (Value) Section ¶

func (v Value) Section() string

func (Value) SetAlignment ¶

func (v Value) SetAlignment(a int)

func (Value) SetAtomicSingleThread ¶

func (v Value) SetAtomicSingleThread(singleThread bool)

func (Value) SetCleanup ¶

func (l Value) SetCleanup(cleanup bool)

func (Value) SetCmpXchgFailureOrdering ¶

func (v Value) SetCmpXchgFailureOrdering(ordering AtomicOrdering)

func (Value) SetCmpXchgSuccessOrdering ¶

func (v Value) SetCmpXchgSuccessOrdering(ordering AtomicOrdering)

func (Value) SetComdat ¶

func (v Value) SetComdat(c Comdat)

func (Value) SetFunctionCallConv ¶

func (v Value) SetFunctionCallConv(cc CallConv)

func (Value) SetGC ¶

func (v Value) SetGC(name string)

func (Value) SetGlobalConstant ¶

func (v Value) SetGlobalConstant(gc bool)

func (Value) SetInitializer ¶

func (v Value) SetInitializer(cv Value)

func (Value) SetInstrParamAlignment ¶

func (v Value) SetInstrParamAlignment(i int, align int)

func (Value) SetInstructionCallConv ¶

func (v Value) SetInstructionCallConv(cc CallConv)

Operations on call sites
func (Value) SetLinkage ¶

func (v Value) SetLinkage(l Linkage)

func (Value) SetMetadata ¶

func (v Value) SetMetadata(kind int, node Metadata)

func (Value) SetName ¶

func (v Value) SetName(name string)

func (Value) SetOperand ¶

func (v Value) SetOperand(i int, op Value)

func (Value) SetOrdering ¶

func (v Value) SetOrdering(ordering AtomicOrdering)

func (Value) SetParamAlignment ¶

func (v Value) SetParamAlignment(align int)

func (Value) SetParentCatchSwitch ¶

func (v Value) SetParentCatchSwitch(catchSwitch Value)

Set the parent catchswitch instruction of a catchpad instruction.

This only works on llvm::CatchPadInst instructions.
func (Value) SetPersonality ¶

func (v Value) SetPersonality(p Value)

func (Value) SetSection ¶

func (v Value) SetSection(str string)

func (Value) SetSubprogram ¶

func (v Value) SetSubprogram(sp Metadata)

func (Value) SetTailCall ¶

func (v Value) SetTailCall(is bool)

func (Value) SetThreadLocal ¶

func (v Value) SetThreadLocal(tl bool)

func (Value) SetUnnamedAddr ¶

func (v Value) SetUnnamedAddr(ua bool)

func (Value) SetVisibility ¶

func (v Value) SetVisibility(vi Visibility)

func (Value) SetVolatile ¶

func (v Value) SetVolatile(volatile bool)

func (Value) String ¶

func (v Value) String() string

Obtain the string value of the instruction. Same as would be printed with Value.Dump() (with two spaces at the start but no newline at the end).
func (Value) Subprogram ¶

func (v Value) Subprogram() (md Metadata)

func (Value) Type ¶

func (v Value) Type() (t Type)

Operations on all values
func (Value) Visibility ¶

func (v Value) Visibility() Visibility

func (Value) ZExtValue ¶

func (v Value) ZExtValue() uint64

type VerifierFailureAction ¶

type VerifierFailureAction C.LLVMVerifierFailureAction

const (
	// verifier will print to stderr and abort()
	AbortProcessAction VerifierFailureAction = C.LLVMAbortProcessAction
	// verifier will print to stderr and return 1
	PrintMessageAction VerifierFailureAction = C.LLVMPrintMessageAction
	// verifier will just return 1
	ReturnStatusAction VerifierFailureAction = C.LLVMReturnStatusAction
)

type Visibility ¶

type Visibility C.LLVMVisibility

const (
	DefaultVisibility   Visibility = C.LLVMDefaultVisibility
	HiddenVisibility    Visibility = C.LLVMHiddenVisibility
	ProtectedVisibility Visibility = C.LLVMProtectedVisibility
)

