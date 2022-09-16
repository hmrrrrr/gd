// Code generated by /internal/packages/generator DO NOT EDIT
package physics

import "github.com/readykit/gd"

type Area2DSpaceOverride int64

const (
	Area2DSpaceOverrideDisabled Area2DSpaceOverride = 0
	Area2DSpaceOverrideCombine Area2DSpaceOverride = 1
	Area2DSpaceOverrideCombineReplace Area2DSpaceOverride = 2
	Area2DSpaceOverrideReplace Area2DSpaceOverride = 3
	Area2DSpaceOverrideReplaceCombine Area2DSpaceOverride = 4
)

type Area3DSpaceOverride int64

const (
	Area3DSpaceOverrideDisabled Area3DSpaceOverride = 0
	Area3DSpaceOverrideCombine Area3DSpaceOverride = 1
	Area3DSpaceOverrideCombineReplace Area3DSpaceOverride = 2
	Area3DSpaceOverrideReplace Area3DSpaceOverride = 3
	Area3DSpaceOverrideReplaceCombine Area3DSpaceOverride = 4
)

type CharacterBody2DMotionMode int64

const (
	CharacterBody2DMotionModeGrounded CharacterBody2DMotionMode = 0
	CharacterBody2DMotionModeFloating CharacterBody2DMotionMode = 1
)

type CharacterBody2DPlatformOnLeave int64

const (
	CharacterBody2DPlatformOnLeaveAddVelocity CharacterBody2DPlatformOnLeave = 0
	CharacterBody2DPlatformOnLeaveAddUpwardVelocity CharacterBody2DPlatformOnLeave = 1
	CharacterBody2DPlatformOnLeaveDoNothing CharacterBody2DPlatformOnLeave = 2
)

type CharacterBody3DMotionMode int64

const (
	CharacterBody3DMotionModeGrounded CharacterBody3DMotionMode = 0
	CharacterBody3DMotionModeFloating CharacterBody3DMotionMode = 1
)

type CharacterBody3DPlatformOnLeave int64

const (
	CharacterBody3DPlatformOnLeaveAddVelocity CharacterBody3DPlatformOnLeave = 0
	CharacterBody3DPlatformOnLeaveAddUpwardVelocity CharacterBody3DPlatformOnLeave = 1
	CharacterBody3DPlatformOnLeaveDoNothing CharacterBody3DPlatformOnLeave = 2
)

type CollisionObject2DDisableMode int64

const (
	CollisionObject2DDisableModeRemove CollisionObject2DDisableMode = 0
	CollisionObject2DDisableModeMakeStatic CollisionObject2DDisableMode = 1
	CollisionObject2DDisableModeKeepActive CollisionObject2DDisableMode = 2
)

type CollisionObject3DDisableMode int64

const (
	CollisionObject3DDisableModeRemove CollisionObject3DDisableMode = 0
	CollisionObject3DDisableModeMakeStatic CollisionObject3DDisableMode = 1
	CollisionObject3DDisableModeKeepActive CollisionObject3DDisableMode = 2
)

type CollisionPolygon2DBuildMode int64

const (
	CollisionPolygon2DBuildSolids CollisionPolygon2DBuildMode = 0
	CollisionPolygon2DBuildSegments CollisionPolygon2DBuildMode = 1
)

type ConeTwistJoint3DParam int64

const (
	ConeTwistJoint3DParamSwingSpan ConeTwistJoint3DParam = 0
	ConeTwistJoint3DParamTwistSpan ConeTwistJoint3DParam = 1
	ConeTwistJoint3DParamBias ConeTwistJoint3DParam = 2
	ConeTwistJoint3DParamSoftness ConeTwistJoint3DParam = 3
	ConeTwistJoint3DParamRelaxation ConeTwistJoint3DParam = 4
	ConeTwistJoint3DParamMax ConeTwistJoint3DParam = 5
)

type Generic6DOFJoint3DParam int64

const (
	Generic6DOFJoint3DParamLinearLowerLimit Generic6DOFJoint3DParam = 0
	Generic6DOFJoint3DParamLinearUpperLimit Generic6DOFJoint3DParam = 1
	Generic6DOFJoint3DParamLinearLimitSoftness Generic6DOFJoint3DParam = 2
	Generic6DOFJoint3DParamLinearRestitution Generic6DOFJoint3DParam = 3
	Generic6DOFJoint3DParamLinearDamping Generic6DOFJoint3DParam = 4
	Generic6DOFJoint3DParamLinearMotorTargetVelocity Generic6DOFJoint3DParam = 5
	Generic6DOFJoint3DParamLinearMotorForceLimit Generic6DOFJoint3DParam = 6
	Generic6DOFJoint3DParamLinearSpringStiffness Generic6DOFJoint3DParam = 7
	Generic6DOFJoint3DParamLinearSpringDamping Generic6DOFJoint3DParam = 8
	Generic6DOFJoint3DParamLinearSpringEquilibriumPoint Generic6DOFJoint3DParam = 9
	Generic6DOFJoint3DParamAngularLowerLimit Generic6DOFJoint3DParam = 10
	Generic6DOFJoint3DParamAngularUpperLimit Generic6DOFJoint3DParam = 11
	Generic6DOFJoint3DParamAngularLimitSoftness Generic6DOFJoint3DParam = 12
	Generic6DOFJoint3DParamAngularDamping Generic6DOFJoint3DParam = 13
	Generic6DOFJoint3DParamAngularRestitution Generic6DOFJoint3DParam = 14
	Generic6DOFJoint3DParamAngularForceLimit Generic6DOFJoint3DParam = 15
	Generic6DOFJoint3DParamAngularErp Generic6DOFJoint3DParam = 16
	Generic6DOFJoint3DParamAngularMotorTargetVelocity Generic6DOFJoint3DParam = 17
	Generic6DOFJoint3DParamAngularMotorForceLimit Generic6DOFJoint3DParam = 18
	Generic6DOFJoint3DParamAngularSpringStiffness Generic6DOFJoint3DParam = 19
	Generic6DOFJoint3DParamAngularSpringDamping Generic6DOFJoint3DParam = 20
	Generic6DOFJoint3DParamAngularSpringEquilibriumPoint Generic6DOFJoint3DParam = 21
	Generic6DOFJoint3DParamMax Generic6DOFJoint3DParam = 22
)

type Generic6DOFJoint3DFlag int64

const (
	Generic6DOFJoint3DFlagEnableLinearLimit Generic6DOFJoint3DFlag = 0
	Generic6DOFJoint3DFlagEnableAngularLimit Generic6DOFJoint3DFlag = 1
	Generic6DOFJoint3DFlagEnableLinearSpring Generic6DOFJoint3DFlag = 3
	Generic6DOFJoint3DFlagEnableAngularSpring Generic6DOFJoint3DFlag = 2
	Generic6DOFJoint3DFlagEnableMotor Generic6DOFJoint3DFlag = 4
	Generic6DOFJoint3DFlagEnableLinearMotor Generic6DOFJoint3DFlag = 5
	Generic6DOFJoint3DFlagMax Generic6DOFJoint3DFlag = 6
)

type HingeJoint3DParam int64

const (
	HingeJoint3DParamBias HingeJoint3DParam = 0
	HingeJoint3DParamLimitUpper HingeJoint3DParam = 1
	HingeJoint3DParamLimitLower HingeJoint3DParam = 2
	HingeJoint3DParamLimitBias HingeJoint3DParam = 3
	HingeJoint3DParamLimitSoftness HingeJoint3DParam = 4
	HingeJoint3DParamLimitRelaxation HingeJoint3DParam = 5
	HingeJoint3DParamMotorTargetVelocity HingeJoint3DParam = 6
	HingeJoint3DParamMotorMaxImpulse HingeJoint3DParam = 7
	HingeJoint3DParamMax HingeJoint3DParam = 8
)

type HingeJoint3DFlag int64

const (
	HingeJoint3DFlagUseLimit HingeJoint3DFlag = 0
	HingeJoint3DFlagEnableMotor HingeJoint3DFlag = 1
	HingeJoint3DFlagMax HingeJoint3DFlag = 2
)

type PhysicalBone3DDampMode int64

const (
	PhysicalBone3DDampModeCombine PhysicalBone3DDampMode = 0
	PhysicalBone3DDampModeReplace PhysicalBone3DDampMode = 1
)

type PhysicalBone3DJointType int64

const (
	PhysicalBone3DJointTypeNone PhysicalBone3DJointType = 0
	PhysicalBone3DJointTypePin PhysicalBone3DJointType = 1
	PhysicalBone3DJointTypeCone PhysicalBone3DJointType = 2
	PhysicalBone3DJointTypeHinge PhysicalBone3DJointType = 3
	PhysicalBone3DJointTypeSlider PhysicalBone3DJointType = 4
	PhysicalBone3DJointType6dof PhysicalBone3DJointType = 5
)

type Server2DSingletonSpaceParameter int64

const (
	Server2DSingletonSpaceParamContactRecycleRadius Server2DSingletonSpaceParameter = 0
	Server2DSingletonSpaceParamContactMaxSeparation Server2DSingletonSpaceParameter = 1
	Server2DSingletonSpaceParamContactMaxAllowedPenetration Server2DSingletonSpaceParameter = 2
	Server2DSingletonSpaceParamContactDefaultBias Server2DSingletonSpaceParameter = 3
	Server2DSingletonSpaceParamBodyLinearVelocitySleepThreshold Server2DSingletonSpaceParameter = 4
	Server2DSingletonSpaceParamBodyAngularVelocitySleepThreshold Server2DSingletonSpaceParameter = 5
	Server2DSingletonSpaceParamBodyTimeToSleep Server2DSingletonSpaceParameter = 6
	Server2DSingletonSpaceParamConstraintDefaultBias Server2DSingletonSpaceParameter = 7
	Server2DSingletonSpaceParamSolverIterations Server2DSingletonSpaceParameter = 8
)

type Server2DSingletonShapeType int64

const (
	Server2DSingletonShapeWorldBoundary Server2DSingletonShapeType = 0
	Server2DSingletonShapeSeparationRay Server2DSingletonShapeType = 1
	Server2DSingletonShapeSegment Server2DSingletonShapeType = 2
	Server2DSingletonShapeCircle Server2DSingletonShapeType = 3
	Server2DSingletonShapeRectangle Server2DSingletonShapeType = 4
	Server2DSingletonShapeCapsule Server2DSingletonShapeType = 5
	Server2DSingletonShapeConvexPolygon Server2DSingletonShapeType = 6
	Server2DSingletonShapeConcavePolygon Server2DSingletonShapeType = 7
	Server2DSingletonShapeCustom Server2DSingletonShapeType = 8
)

type Server2DSingletonAreaParameter int64

const (
	Server2DSingletonAreaParamGravityOverrideMode Server2DSingletonAreaParameter = 0
	Server2DSingletonAreaParamGravity Server2DSingletonAreaParameter = 1
	Server2DSingletonAreaParamGravityVector Server2DSingletonAreaParameter = 2
	Server2DSingletonAreaParamGravityIsPoint Server2DSingletonAreaParameter = 3
	Server2DSingletonAreaParamGravityDistanceScale Server2DSingletonAreaParameter = 4
	Server2DSingletonAreaParamGravityPointAttenuation Server2DSingletonAreaParameter = 5
	Server2DSingletonAreaParamLinearDampOverrideMode Server2DSingletonAreaParameter = 6
	Server2DSingletonAreaParamLinearDamp Server2DSingletonAreaParameter = 7
	Server2DSingletonAreaParamAngularDampOverrideMode Server2DSingletonAreaParameter = 8
	Server2DSingletonAreaParamAngularDamp Server2DSingletonAreaParameter = 9
	Server2DSingletonAreaParamPriority Server2DSingletonAreaParameter = 10
)

type Server2DSingletonAreaSpaceOverrideMode int64

const (
	Server2DSingletonAreaSpaceOverrideDisabled Server2DSingletonAreaSpaceOverrideMode = 0
	Server2DSingletonAreaSpaceOverrideCombine Server2DSingletonAreaSpaceOverrideMode = 1
	Server2DSingletonAreaSpaceOverrideCombineReplace Server2DSingletonAreaSpaceOverrideMode = 2
	Server2DSingletonAreaSpaceOverrideReplace Server2DSingletonAreaSpaceOverrideMode = 3
	Server2DSingletonAreaSpaceOverrideReplaceCombine Server2DSingletonAreaSpaceOverrideMode = 4
)

type Server2DSingletonBodyMode int64

const (
	Server2DSingletonBodyModeStatic Server2DSingletonBodyMode = 0
	Server2DSingletonBodyModeKinematic Server2DSingletonBodyMode = 1
	Server2DSingletonBodyModeRigid Server2DSingletonBodyMode = 2
	Server2DSingletonBodyModeRigidLinear Server2DSingletonBodyMode = 3
)

type Server2DSingletonBodyParameter int64

const (
	Server2DSingletonBodyParamBounce Server2DSingletonBodyParameter = 0
	Server2DSingletonBodyParamFriction Server2DSingletonBodyParameter = 1
	Server2DSingletonBodyParamMass Server2DSingletonBodyParameter = 2
	Server2DSingletonBodyParamInertia Server2DSingletonBodyParameter = 3
	Server2DSingletonBodyParamCenterOfMass Server2DSingletonBodyParameter = 4
	Server2DSingletonBodyParamGravityScale Server2DSingletonBodyParameter = 5
	Server2DSingletonBodyParamLinearDampMode Server2DSingletonBodyParameter = 6
	Server2DSingletonBodyParamAngularDampMode Server2DSingletonBodyParameter = 7
	Server2DSingletonBodyParamLinearDamp Server2DSingletonBodyParameter = 8
	Server2DSingletonBodyParamAngularDamp Server2DSingletonBodyParameter = 9
	Server2DSingletonBodyParamMax Server2DSingletonBodyParameter = 10
)

type Server2DSingletonBodyDampMode int64

const (
	Server2DSingletonBodyDampModeCombine Server2DSingletonBodyDampMode = 0
	Server2DSingletonBodyDampModeReplace Server2DSingletonBodyDampMode = 1
)

type Server2DSingletonBodyState int64

const (
	Server2DSingletonBodyStateTransform Server2DSingletonBodyState = 0
	Server2DSingletonBodyStateLinearVelocity Server2DSingletonBodyState = 1
	Server2DSingletonBodyStateAngularVelocity Server2DSingletonBodyState = 2
	Server2DSingletonBodyStateSleeping Server2DSingletonBodyState = 3
	Server2DSingletonBodyStateCanSleep Server2DSingletonBodyState = 4
)

type Server2DSingletonJointType int64

const (
	Server2DSingletonJointTypePin Server2DSingletonJointType = 0
	Server2DSingletonJointTypeGroove Server2DSingletonJointType = 1
	Server2DSingletonJointTypeDampedSpring Server2DSingletonJointType = 2
	Server2DSingletonJointTypeMax Server2DSingletonJointType = 3
)

type Server2DSingletonJointParam int64

const (
	Server2DSingletonJointParamBias Server2DSingletonJointParam = 0
	Server2DSingletonJointParamMaxBias Server2DSingletonJointParam = 1
	Server2DSingletonJointParamMaxForce Server2DSingletonJointParam = 2
)

type Server2DSingletonDampedSpringParam int64

const (
	Server2DSingletonDampedSpringRestLength Server2DSingletonDampedSpringParam = 0
	Server2DSingletonDampedSpringStiffness Server2DSingletonDampedSpringParam = 1
	Server2DSingletonDampedSpringDamping Server2DSingletonDampedSpringParam = 2
)

type Server2DSingletonCCDMode int64

const (
	Server2DSingletonCcdModeDisabled Server2DSingletonCCDMode = 0
	Server2DSingletonCcdModeCastRay Server2DSingletonCCDMode = 1
	Server2DSingletonCcdModeCastShape Server2DSingletonCCDMode = 2
)

type Server2DSingletonAreaBodyStatus int64

const (
	Server2DSingletonAreaBodyAdded Server2DSingletonAreaBodyStatus = 0
	Server2DSingletonAreaBodyRemoved Server2DSingletonAreaBodyStatus = 1
)

type Server2DSingletonProcessInfo int64

const (
	Server2DSingletonInfoActiveObjects Server2DSingletonProcessInfo = 0
	Server2DSingletonInfoCollisionPairs Server2DSingletonProcessInfo = 1
	Server2DSingletonInfoIslandCount Server2DSingletonProcessInfo = 2
)

type Server3DSingletonJointType int64

const (
	Server3DSingletonJointTypePin Server3DSingletonJointType = 0
	Server3DSingletonJointTypeHinge Server3DSingletonJointType = 1
	Server3DSingletonJointTypeSlider Server3DSingletonJointType = 2
	Server3DSingletonJointTypeConeTwist Server3DSingletonJointType = 3
	Server3DSingletonJointType6dof Server3DSingletonJointType = 4
	Server3DSingletonJointTypeMax Server3DSingletonJointType = 5
)

type Server3DSingletonPinJointParam int64

const (
	Server3DSingletonPinJointBias Server3DSingletonPinJointParam = 0
	Server3DSingletonPinJointDamping Server3DSingletonPinJointParam = 1
	Server3DSingletonPinJointImpulseClamp Server3DSingletonPinJointParam = 2
)

type Server3DSingletonHingeJointParam int64

const (
	Server3DSingletonHingeJointBias Server3DSingletonHingeJointParam = 0
	Server3DSingletonHingeJointLimitUpper Server3DSingletonHingeJointParam = 1
	Server3DSingletonHingeJointLimitLower Server3DSingletonHingeJointParam = 2
	Server3DSingletonHingeJointLimitBias Server3DSingletonHingeJointParam = 3
	Server3DSingletonHingeJointLimitSoftness Server3DSingletonHingeJointParam = 4
	Server3DSingletonHingeJointLimitRelaxation Server3DSingletonHingeJointParam = 5
	Server3DSingletonHingeJointMotorTargetVelocity Server3DSingletonHingeJointParam = 6
	Server3DSingletonHingeJointMotorMaxImpulse Server3DSingletonHingeJointParam = 7
)

type Server3DSingletonHingeJointFlag int64

const (
	Server3DSingletonHingeJointFlagUseLimit Server3DSingletonHingeJointFlag = 0
	Server3DSingletonHingeJointFlagEnableMotor Server3DSingletonHingeJointFlag = 1
)

type Server3DSingletonSliderJointParam int64

const (
	Server3DSingletonSliderJointLinearLimitUpper Server3DSingletonSliderJointParam = 0
	Server3DSingletonSliderJointLinearLimitLower Server3DSingletonSliderJointParam = 1
	Server3DSingletonSliderJointLinearLimitSoftness Server3DSingletonSliderJointParam = 2
	Server3DSingletonSliderJointLinearLimitRestitution Server3DSingletonSliderJointParam = 3
	Server3DSingletonSliderJointLinearLimitDamping Server3DSingletonSliderJointParam = 4
	Server3DSingletonSliderJointLinearMotionSoftness Server3DSingletonSliderJointParam = 5
	Server3DSingletonSliderJointLinearMotionRestitution Server3DSingletonSliderJointParam = 6
	Server3DSingletonSliderJointLinearMotionDamping Server3DSingletonSliderJointParam = 7
	Server3DSingletonSliderJointLinearOrthogonalSoftness Server3DSingletonSliderJointParam = 8
	Server3DSingletonSliderJointLinearOrthogonalRestitution Server3DSingletonSliderJointParam = 9
	Server3DSingletonSliderJointLinearOrthogonalDamping Server3DSingletonSliderJointParam = 10
	Server3DSingletonSliderJointAngularLimitUpper Server3DSingletonSliderJointParam = 11
	Server3DSingletonSliderJointAngularLimitLower Server3DSingletonSliderJointParam = 12
	Server3DSingletonSliderJointAngularLimitSoftness Server3DSingletonSliderJointParam = 13
	Server3DSingletonSliderJointAngularLimitRestitution Server3DSingletonSliderJointParam = 14
	Server3DSingletonSliderJointAngularLimitDamping Server3DSingletonSliderJointParam = 15
	Server3DSingletonSliderJointAngularMotionSoftness Server3DSingletonSliderJointParam = 16
	Server3DSingletonSliderJointAngularMotionRestitution Server3DSingletonSliderJointParam = 17
	Server3DSingletonSliderJointAngularMotionDamping Server3DSingletonSliderJointParam = 18
	Server3DSingletonSliderJointAngularOrthogonalSoftness Server3DSingletonSliderJointParam = 19
	Server3DSingletonSliderJointAngularOrthogonalRestitution Server3DSingletonSliderJointParam = 20
	Server3DSingletonSliderJointAngularOrthogonalDamping Server3DSingletonSliderJointParam = 21
	Server3DSingletonSliderJointMax Server3DSingletonSliderJointParam = 22
)

type Server3DSingletonConeTwistJointParam int64

const (
	Server3DSingletonConeTwistJointSwingSpan Server3DSingletonConeTwistJointParam = 0
	Server3DSingletonConeTwistJointTwistSpan Server3DSingletonConeTwistJointParam = 1
	Server3DSingletonConeTwistJointBias Server3DSingletonConeTwistJointParam = 2
	Server3DSingletonConeTwistJointSoftness Server3DSingletonConeTwistJointParam = 3
	Server3DSingletonConeTwistJointRelaxation Server3DSingletonConeTwistJointParam = 4
)

type Server3DSingletonG6DOFJointAxisParam int64

const (
	Server3DSingletonG6dofJointLinearLowerLimit Server3DSingletonG6DOFJointAxisParam = 0
	Server3DSingletonG6dofJointLinearUpperLimit Server3DSingletonG6DOFJointAxisParam = 1
	Server3DSingletonG6dofJointLinearLimitSoftness Server3DSingletonG6DOFJointAxisParam = 2
	Server3DSingletonG6dofJointLinearRestitution Server3DSingletonG6DOFJointAxisParam = 3
	Server3DSingletonG6dofJointLinearDamping Server3DSingletonG6DOFJointAxisParam = 4
	Server3DSingletonG6dofJointLinearMotorTargetVelocity Server3DSingletonG6DOFJointAxisParam = 5
	Server3DSingletonG6dofJointLinearMotorForceLimit Server3DSingletonG6DOFJointAxisParam = 6
	Server3DSingletonG6dofJointAngularLowerLimit Server3DSingletonG6DOFJointAxisParam = 10
	Server3DSingletonG6dofJointAngularUpperLimit Server3DSingletonG6DOFJointAxisParam = 11
	Server3DSingletonG6dofJointAngularLimitSoftness Server3DSingletonG6DOFJointAxisParam = 12
	Server3DSingletonG6dofJointAngularDamping Server3DSingletonG6DOFJointAxisParam = 13
	Server3DSingletonG6dofJointAngularRestitution Server3DSingletonG6DOFJointAxisParam = 14
	Server3DSingletonG6dofJointAngularForceLimit Server3DSingletonG6DOFJointAxisParam = 15
	Server3DSingletonG6dofJointAngularErp Server3DSingletonG6DOFJointAxisParam = 16
	Server3DSingletonG6dofJointAngularMotorTargetVelocity Server3DSingletonG6DOFJointAxisParam = 17
	Server3DSingletonG6dofJointAngularMotorForceLimit Server3DSingletonG6DOFJointAxisParam = 18
)

type Server3DSingletonG6DOFJointAxisFlag int64

const (
	Server3DSingletonG6dofJointFlagEnableLinearLimit Server3DSingletonG6DOFJointAxisFlag = 0
	Server3DSingletonG6dofJointFlagEnableAngularLimit Server3DSingletonG6DOFJointAxisFlag = 1
	Server3DSingletonG6dofJointFlagEnableMotor Server3DSingletonG6DOFJointAxisFlag = 4
	Server3DSingletonG6dofJointFlagEnableLinearMotor Server3DSingletonG6DOFJointAxisFlag = 5
)

type Server3DSingletonShapeType int64

const (
	Server3DSingletonShapeWorldBoundary Server3DSingletonShapeType = 0
	Server3DSingletonShapeSeparationRay Server3DSingletonShapeType = 1
	Server3DSingletonShapeSphere Server3DSingletonShapeType = 2
	Server3DSingletonShapeBox Server3DSingletonShapeType = 3
	Server3DSingletonShapeCapsule Server3DSingletonShapeType = 4
	Server3DSingletonShapeCylinder Server3DSingletonShapeType = 5
	Server3DSingletonShapeConvexPolygon Server3DSingletonShapeType = 6
	Server3DSingletonShapeConcavePolygon Server3DSingletonShapeType = 7
	Server3DSingletonShapeHeightmap Server3DSingletonShapeType = 8
	Server3DSingletonShapeSoftBody Server3DSingletonShapeType = 9
	Server3DSingletonShapeCustom Server3DSingletonShapeType = 10
)

type Server3DSingletonAreaParameter int64

const (
	Server3DSingletonAreaParamGravityOverrideMode Server3DSingletonAreaParameter = 0
	Server3DSingletonAreaParamGravity Server3DSingletonAreaParameter = 1
	Server3DSingletonAreaParamGravityVector Server3DSingletonAreaParameter = 2
	Server3DSingletonAreaParamGravityIsPoint Server3DSingletonAreaParameter = 3
	Server3DSingletonAreaParamGravityDistanceScale Server3DSingletonAreaParameter = 4
	Server3DSingletonAreaParamGravityPointAttenuation Server3DSingletonAreaParameter = 5
	Server3DSingletonAreaParamLinearDampOverrideMode Server3DSingletonAreaParameter = 6
	Server3DSingletonAreaParamLinearDamp Server3DSingletonAreaParameter = 7
	Server3DSingletonAreaParamAngularDampOverrideMode Server3DSingletonAreaParameter = 8
	Server3DSingletonAreaParamAngularDamp Server3DSingletonAreaParameter = 9
	Server3DSingletonAreaParamPriority Server3DSingletonAreaParameter = 10
	Server3DSingletonAreaParamWindForceMagnitude Server3DSingletonAreaParameter = 11
	Server3DSingletonAreaParamWindSource Server3DSingletonAreaParameter = 12
	Server3DSingletonAreaParamWindDirection Server3DSingletonAreaParameter = 13
	Server3DSingletonAreaParamWindAttenuationFactor Server3DSingletonAreaParameter = 14
)

type Server3DSingletonAreaSpaceOverrideMode int64

const (
	Server3DSingletonAreaSpaceOverrideDisabled Server3DSingletonAreaSpaceOverrideMode = 0
	Server3DSingletonAreaSpaceOverrideCombine Server3DSingletonAreaSpaceOverrideMode = 1
	Server3DSingletonAreaSpaceOverrideCombineReplace Server3DSingletonAreaSpaceOverrideMode = 2
	Server3DSingletonAreaSpaceOverrideReplace Server3DSingletonAreaSpaceOverrideMode = 3
	Server3DSingletonAreaSpaceOverrideReplaceCombine Server3DSingletonAreaSpaceOverrideMode = 4
)

type Server3DSingletonBodyMode int64

const (
	Server3DSingletonBodyModeStatic Server3DSingletonBodyMode = 0
	Server3DSingletonBodyModeKinematic Server3DSingletonBodyMode = 1
	Server3DSingletonBodyModeRigid Server3DSingletonBodyMode = 2
	Server3DSingletonBodyModeRigidLinear Server3DSingletonBodyMode = 3
)

type Server3DSingletonBodyParameter int64

const (
	Server3DSingletonBodyParamBounce Server3DSingletonBodyParameter = 0
	Server3DSingletonBodyParamFriction Server3DSingletonBodyParameter = 1
	Server3DSingletonBodyParamMass Server3DSingletonBodyParameter = 2
	Server3DSingletonBodyParamInertia Server3DSingletonBodyParameter = 3
	Server3DSingletonBodyParamCenterOfMass Server3DSingletonBodyParameter = 4
	Server3DSingletonBodyParamGravityScale Server3DSingletonBodyParameter = 5
	Server3DSingletonBodyParamLinearDampMode Server3DSingletonBodyParameter = 6
	Server3DSingletonBodyParamAngularDampMode Server3DSingletonBodyParameter = 7
	Server3DSingletonBodyParamLinearDamp Server3DSingletonBodyParameter = 8
	Server3DSingletonBodyParamAngularDamp Server3DSingletonBodyParameter = 9
	Server3DSingletonBodyParamMax Server3DSingletonBodyParameter = 10
)

type Server3DSingletonBodyDampMode int64

const (
	Server3DSingletonBodyDampModeCombine Server3DSingletonBodyDampMode = 0
	Server3DSingletonBodyDampModeReplace Server3DSingletonBodyDampMode = 1
)

type Server3DSingletonBodyState int64

const (
	Server3DSingletonBodyStateTransform Server3DSingletonBodyState = 0
	Server3DSingletonBodyStateLinearVelocity Server3DSingletonBodyState = 1
	Server3DSingletonBodyStateAngularVelocity Server3DSingletonBodyState = 2
	Server3DSingletonBodyStateSleeping Server3DSingletonBodyState = 3
	Server3DSingletonBodyStateCanSleep Server3DSingletonBodyState = 4
)

type Server3DSingletonAreaBodyStatus int64

const (
	Server3DSingletonAreaBodyAdded Server3DSingletonAreaBodyStatus = 0
	Server3DSingletonAreaBodyRemoved Server3DSingletonAreaBodyStatus = 1
)

type Server3DSingletonProcessInfo int64

const (
	Server3DSingletonInfoActiveObjects Server3DSingletonProcessInfo = 0
	Server3DSingletonInfoCollisionPairs Server3DSingletonProcessInfo = 1
	Server3DSingletonInfoIslandCount Server3DSingletonProcessInfo = 2
)

type Server3DSingletonSpaceParameter int64

const (
	Server3DSingletonSpaceParamContactRecycleRadius Server3DSingletonSpaceParameter = 0
	Server3DSingletonSpaceParamContactMaxSeparation Server3DSingletonSpaceParameter = 1
	Server3DSingletonSpaceParamContactMaxAllowedPenetration Server3DSingletonSpaceParameter = 2
	Server3DSingletonSpaceParamContactDefaultBias Server3DSingletonSpaceParameter = 3
	Server3DSingletonSpaceParamBodyLinearVelocitySleepThreshold Server3DSingletonSpaceParameter = 4
	Server3DSingletonSpaceParamBodyAngularVelocitySleepThreshold Server3DSingletonSpaceParameter = 5
	Server3DSingletonSpaceParamBodyTimeToSleep Server3DSingletonSpaceParameter = 6
	Server3DSingletonSpaceParamSolverIterations Server3DSingletonSpaceParameter = 7
)

type Server3DSingletonBodyAxis int64

const (
	Server3DSingletonBodyAxisLinearX Server3DSingletonBodyAxis = 1
	Server3DSingletonBodyAxisLinearY Server3DSingletonBodyAxis = 2
	Server3DSingletonBodyAxisLinearZ Server3DSingletonBodyAxis = 4
	Server3DSingletonBodyAxisAngularX Server3DSingletonBodyAxis = 8
	Server3DSingletonBodyAxisAngularY Server3DSingletonBodyAxis = 16
	Server3DSingletonBodyAxisAngularZ Server3DSingletonBodyAxis = 32
)

type PinJoint3DParam int64

const (
	PinJoint3DParamBias PinJoint3DParam = 0
	PinJoint3DParamDamping PinJoint3DParam = 1
	PinJoint3DParamImpulseClamp PinJoint3DParam = 2
)

type RigidBody2DFreezeMode int64

const (
	RigidBody2DFreezeModeStatic RigidBody2DFreezeMode = 0
	RigidBody2DFreezeModeKinematic RigidBody2DFreezeMode = 1
)

type RigidBody2DCenterOfMassMode int64

const (
	RigidBody2DCenterOfMassModeAuto RigidBody2DCenterOfMassMode = 0
	RigidBody2DCenterOfMassModeCustom RigidBody2DCenterOfMassMode = 1
)

type RigidBody2DDampMode int64

const (
	RigidBody2DDampModeCombine RigidBody2DDampMode = 0
	RigidBody2DDampModeReplace RigidBody2DDampMode = 1
)

type RigidBody2DCCDMode int64

const (
	RigidBody2DCcdModeDisabled RigidBody2DCCDMode = 0
	RigidBody2DCcdModeCastRay RigidBody2DCCDMode = 1
	RigidBody2DCcdModeCastShape RigidBody2DCCDMode = 2
)

type RigidBody3DFreezeMode int64

const (
	RigidBody3DFreezeModeStatic RigidBody3DFreezeMode = 0
	RigidBody3DFreezeModeKinematic RigidBody3DFreezeMode = 1
)

type RigidBody3DCenterOfMassMode int64

const (
	RigidBody3DCenterOfMassModeAuto RigidBody3DCenterOfMassMode = 0
	RigidBody3DCenterOfMassModeCustom RigidBody3DCenterOfMassMode = 1
)

type RigidBody3DDampMode int64

const (
	RigidBody3DDampModeCombine RigidBody3DDampMode = 0
	RigidBody3DDampModeReplace RigidBody3DDampMode = 1
)

type SliderJoint3DParam int64

const (
	SliderJoint3DParamLinearLimitUpper SliderJoint3DParam = 0
	SliderJoint3DParamLinearLimitLower SliderJoint3DParam = 1
	SliderJoint3DParamLinearLimitSoftness SliderJoint3DParam = 2
	SliderJoint3DParamLinearLimitRestitution SliderJoint3DParam = 3
	SliderJoint3DParamLinearLimitDamping SliderJoint3DParam = 4
	SliderJoint3DParamLinearMotionSoftness SliderJoint3DParam = 5
	SliderJoint3DParamLinearMotionRestitution SliderJoint3DParam = 6
	SliderJoint3DParamLinearMotionDamping SliderJoint3DParam = 7
	SliderJoint3DParamLinearOrthogonalSoftness SliderJoint3DParam = 8
	SliderJoint3DParamLinearOrthogonalRestitution SliderJoint3DParam = 9
	SliderJoint3DParamLinearOrthogonalDamping SliderJoint3DParam = 10
	SliderJoint3DParamAngularLimitUpper SliderJoint3DParam = 11
	SliderJoint3DParamAngularLimitLower SliderJoint3DParam = 12
	SliderJoint3DParamAngularLimitSoftness SliderJoint3DParam = 13
	SliderJoint3DParamAngularLimitRestitution SliderJoint3DParam = 14
	SliderJoint3DParamAngularLimitDamping SliderJoint3DParam = 15
	SliderJoint3DParamAngularMotionSoftness SliderJoint3DParam = 16
	SliderJoint3DParamAngularMotionRestitution SliderJoint3DParam = 17
	SliderJoint3DParamAngularMotionDamping SliderJoint3DParam = 18
	SliderJoint3DParamAngularOrthogonalSoftness SliderJoint3DParam = 19
	SliderJoint3DParamAngularOrthogonalRestitution SliderJoint3DParam = 20
	SliderJoint3DParamAngularOrthogonalDamping SliderJoint3DParam = 21
	SliderJoint3DParamMax SliderJoint3DParam = 22
)

type SoftBody3DDisableMode int64

const (
	SoftBody3DDisableModeRemove SoftBody3DDisableMode = 0
	SoftBody3DDisableModeKeepActive SoftBody3DDisableMode = 1
)

type Area2D = gd.Area2D

type Area3D = gd.Area3D

type BoxShape3D = gd.BoxShape3D

type CapsuleShape2D = gd.CapsuleShape2D

type CapsuleShape3D = gd.CapsuleShape3D

type CharacterBody2D = gd.CharacterBody2D

type CharacterBody3D = gd.CharacterBody3D

type CircleShape2D = gd.CircleShape2D

type CollisionObject2D = gd.CollisionObject2D

type CollisionObject3D = gd.CollisionObject3D

type CollisionPolygon2D = gd.CollisionPolygon2D

type CollisionPolygon3D = gd.CollisionPolygon3D

type CollisionShape2D = gd.CollisionShape2D

type CollisionShape3D = gd.CollisionShape3D

type ConcavePolygonShape2D = gd.ConcavePolygonShape2D

type ConcavePolygonShape3D = gd.ConcavePolygonShape3D

type ConeTwistJoint3D = gd.ConeTwistJoint3D

type ConvexPolygonShape2D = gd.ConvexPolygonShape2D

type ConvexPolygonShape3D = gd.ConvexPolygonShape3D

type CylinderShape3D = gd.CylinderShape3D

type DampedSpringJoint2D = gd.DampedSpringJoint2D

type Generic6DOFJoint3D = gd.Generic6DOFJoint3D

type GodotServer2D = gd.GodotPhysicsServer2D

type GodotServer3D = gd.GodotPhysicsServer3D

type GrooveJoint2D = gd.GrooveJoint2D

type HeightMapShape3D = gd.HeightMapShape3D

type HingeJoint3D = gd.HingeJoint3D

type Joint2D = gd.Joint2D

type Joint3D = gd.Joint3D

type KinematicCollision2D = gd.KinematicCollision2D

type KinematicCollision3D = gd.KinematicCollision3D

type PhysicalBone2D = gd.PhysicalBone2D

type PhysicalBone3D = gd.PhysicalBone3D

type PhysicalSkyMaterial = gd.PhysicalSkyMaterial

type Body2D = gd.PhysicsBody2D

type Body3D = gd.PhysicsBody3D

type DirectBodyState2D = gd.PhysicsDirectBodyState2D

type DirectBodyState3D = gd.PhysicsDirectBodyState3D

type DirectBodyState3DExtension = gd.PhysicsDirectBodyState3DExtension

type DirectSpaceState2D = gd.PhysicsDirectSpaceState2D

type DirectSpaceState3D = gd.PhysicsDirectSpaceState3D

type DirectSpaceState3DExtension = gd.PhysicsDirectSpaceState3DExtension

type Material = gd.PhysicsMaterial

type PointQueryParameters2D = gd.PhysicsPointQueryParameters2D

type PointQueryParameters3D = gd.PhysicsPointQueryParameters3D

type RayQueryParameters2D = gd.PhysicsRayQueryParameters2D

type RayQueryParameters3D = gd.PhysicsRayQueryParameters3D

var Server2D Server2DSingleton

type Server2DSingleton = gd.PhysicsServer2DSingleton

var Server3D Server3DSingleton

type Server3DSingleton = gd.PhysicsServer3DSingleton

type Server3DExtension = gd.PhysicsServer3DExtension

type Server3DRenderingServerHandler = gd.PhysicsServer3DRenderingServerHandler

type ShapeQueryParameters2D = gd.PhysicsShapeQueryParameters2D

type ShapeQueryParameters3D = gd.PhysicsShapeQueryParameters3D

type TestMotionParameters2D = gd.PhysicsTestMotionParameters2D

type TestMotionParameters3D = gd.PhysicsTestMotionParameters3D

type TestMotionResult2D = gd.PhysicsTestMotionResult2D

type TestMotionResult3D = gd.PhysicsTestMotionResult3D

type PinJoint2D = gd.PinJoint2D

type PinJoint3D = gd.PinJoint3D

type RayCast2D = gd.RayCast2D

type RayCast3D = gd.RayCast3D

type RectangleShape2D = gd.RectangleShape2D

type RigidBody2D = gd.RigidBody2D

type RigidBody3D = gd.RigidBody3D

type SegmentShape2D = gd.SegmentShape2D

type SeparationRayShape2D = gd.SeparationRayShape2D

type SeparationRayShape3D = gd.SeparationRayShape3D

type Shape2D = gd.Shape2D

type Shape3D = gd.Shape3D

type ShapeCast2D = gd.ShapeCast2D

type ShapeCast3D = gd.ShapeCast3D

type SliderJoint3D = gd.SliderJoint3D

type SoftBody3D = gd.SoftBody3D

type SphereShape3D = gd.SphereShape3D

type SpringArm3D = gd.SpringArm3D

type StaticBody2D = gd.StaticBody2D

type StaticBody3D = gd.StaticBody3D

type VehicleBody3D = gd.VehicleBody3D

type VehicleWheel3D = gd.VehicleWheel3D

type WorldBoundaryShape2D = gd.WorldBoundaryShape2D

type WorldBoundaryShape3D = gd.WorldBoundaryShape3D
